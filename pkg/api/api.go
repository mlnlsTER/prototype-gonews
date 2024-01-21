package api

import (
	"GoNews/pkg/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// API - API object.
type API struct {
	db     storage.Interface
	router *mux.Router
}

// Constructor creates a new API object.
func New(db storage.Interface) *API {
	api := API{
		db: db, router: mux.NewRouter(),
	}
	api.endpoints()
	return &api
}

// Registering API endpoints.
func (api *API) endpoints() {
	// get the latest news
	api.router.HandleFunc("/news/{n}", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)
	// web application
	api.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("cmd/server/webapp"))))
}

// Receive request router.
// Required to pass the router to the web server.
func (api *API) Router() *mux.Router {
	return api.router
}

// Handler for the latest news.
func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	s := mux.Vars(r)["n"]
	n, _ := strconv.Atoi(s)
	posts, err := api.db.Posts(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posts)
}
