package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/rss"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/postgres"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

type config struct {
	URLS   []string `json:"rss"`
	Period int      `json:"request_period"`
}

func main() {
	var err error
	var srv server

	// Initialize PostgreSQL server storage.
	srv.db, err = postgres.New("postgres://postgres:8952@localhost:5432/posts")
	if err != nil {
		log.Fatal(err)
	}

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)
	c, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var conf config
	err = json.Unmarshal(c, &conf)
	if err != nil {
		log.Fatal(err)
	}

	chPosts := make(chan []storage.Post)
	chErrors := make(chan error)

	for _, url := range conf.URLS {
		go parseURL(url, chPosts, chErrors, conf.Period)
	}

	go func() {
		for posts := range chPosts {
			for _, post := range posts {
				srv.db.AddPost(post)
			}

		}
	}()

	go func() {
		for err = range chErrors {
			log.Println(err)
		}
	}()

	err = http.ListenAndServe(":8080", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}
}

func parseURL(url string, chPosts chan []storage.Post, chErrors chan error, peroid int) {
	posts, err := rss.Parse(url)
	if err != nil {
		chErrors <- err
		return
	}
	chPosts <- posts
	time.Sleep(time.Duration(peroid) * time.Minute)
}
