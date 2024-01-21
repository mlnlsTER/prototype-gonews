package memdb

import (
	"GoNews/pkg/storage"
)

// Хранилище данных.
type DB []storage.Post

func (db *DB) Posts() ([]storage.Post, error) {
	return *db, nil
}
func (db *DB) AddPost(post storage.Post) error {
	*db = append(*db, post)
	return nil
}
