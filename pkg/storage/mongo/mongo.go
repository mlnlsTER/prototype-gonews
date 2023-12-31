package mongo

import (
	"GoNews/pkg/storage"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "data"  // имя учебной БД
	collectionName = "posts" // имя коллекции в учебной БД
)

type Storage struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func New(mongoDBURL string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(mongoDBURL)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, err
	}
	// не забываем закрывать ресурсы
	defer client.Disconnect(context.Background())
	// проверка связи с БД
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(databaseName)
	collection := db.Collection(collectionName)
	return &Storage{
		client:     client,
		database:   db,
		collection: collection,
	}, nil
}

// Posts возвращает все документы из БД.
func (mdb *Storage) Posts() ([]storage.Post, error) {
	filter := bson.D{}
	cur, err := mdb.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var posts []storage.Post
	for cur.Next(context.Background()) {
		var post storage.Post
		if err := cur.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, cur.Err()
}

// AddPosts вставляет в БД массив документов.
func (mdb *Storage) AddPost(post storage.Post) error {
	_, err := mdb.collection.InsertOne(context.Background(), post)
	if err != nil {
		return err
	}
	return nil
}

func (mdb *Storage) UpdatePost(post storage.Post) error {
	filter := bson.M{"_id": post.ID}
	update := bson.M{"$set": post}

	_, err := mdb.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (mdb *Storage) DeletePost(post storage.Post) error {
	filter := bson.M{"_id": post.ID}
	_, err := mdb.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
