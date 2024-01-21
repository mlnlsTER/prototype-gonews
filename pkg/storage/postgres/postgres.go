package postgres

import (
	"GoNews/pkg/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Data storage.
type Storage struct {
	db *pgxpool.Pool
}

// Constructor creates a new Storage object.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Posts returns the last n publications from the database.
func (s *Storage) Posts(n int) ([]storage.Post, error) {
	if n == 0 {
		n = 6
	}
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			title,
			content,
			published_at,
			link
		FROM posts
		ORDER BY id DESC
		LIMIT $1;
	`, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PubTime,
			&p.Link,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)

	}
	return posts, rows.Err()
}

// AddPosts creates a new publications in the database.
func (s *Storage) AddPosts(posts []storage.Post) error {
	for _, post := range posts {
		_, err := s.db.Exec(context.Background(), `
		INSERT INTO posts(title, content, published_at, link)
		VALUES ($1, $2, $3, $4)`,
			post.Title,
			post.Content,
			post.PubTime,
			post.Link,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
