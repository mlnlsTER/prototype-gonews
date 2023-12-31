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
func (s *Storage) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			author_id,
			title,
			content,
			created_at,
			published_at
		FROM posts
		ORDER BY id;
	`)
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
			&p.AuthorID,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
			&p.PublishedAt,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, p)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return posts, rows.Err()
}

// AddPost создаёт новую публикацию
func (s *Storage) AddPost(p storage.Post) error {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO posts (title, content)
		VALUES ($1, $2) RETURNING id;
		`,
		p.Title,
		p.Content,
	).Scan(&id)
	return err
}

// Обновлять публикацию по id
func (s *Storage) UpdatePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	UPDATE posts
	SET 
		title = $1,
		content = $2,
	WHERE id = $3;
`, post.Title, post.Content, post.ID)
	return err
}

// Удалять публикацию по id
func (s *Storage) DeletePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
        DELETE FROM posts
        WHERE id = $1;
    `, post.ID)
	return err
}

var posts = []storage.Post{
	{
		ID:      1,
		Title:   "Effective Go",
		Content: "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
	},
	{
		ID:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
	},
}
