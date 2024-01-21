package postgres

import (
	"GoNews/pkg/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
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

// Posts возвращает список публикаций из БД.
func (s *Storage) Posts(n int) ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			title,
			content,
			published_at,
			link
		FROM posts
		ORDER BY id
		LIMIT $1;
	`, n,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []storage.Post
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
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
		// добавление переменной в массив результатов
		posts = append(posts, p)

	}
	return posts, rows.Err()
}

// AddPost создаёт новую публикацию
func (s *Storage) AddPost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO posts (title, content, published_at, link)
		VALUES ($1, $2, $3, $4);
	`, p.Title, p.Content, p.PubTime, p.Link)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// Удалять публикацию по id
func (s *Storage) DeletePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
        DELETE FROM posts
        WHERE id = $1;
    `, post.ID)
	return err
}
