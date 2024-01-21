package postgres

import (
	"GoNews/pkg/storage"
	"testing"
)

func TestNew(t *testing.T) {
	constr := "postgres://postgres:8952@localhost:5432/posts"
	_, err := New(constr)
	if err != nil {
		t.Errorf("Failed to create new storage: %v", err)
	}
}

func TestStorage_AddPosts(t *testing.T) {
	constr := "postgres://postgres:8952@localhost:5432/posts"
	s, err := New(constr)
	posts := []storage.Post{
		{
			Title:   "Unit Test Task",
			Content: "Task Content",
			PubTime: 1631234567,
			Link:    "https://example.com/post1",
		},
		{
			Title:   "Unit Test Task 2",
			Content: "Task Content 2",
			PubTime: 1631234568,
			Link:    "https://example.com/post2",
		},
	}
	id := s.AddPosts(posts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создана задача с id:", id)
}
