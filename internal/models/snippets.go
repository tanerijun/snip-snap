package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES (@title, @content, NOW(), NOW() + make_interval(days => @expires))
	RETURNING id`
	args := pgx.NamedArgs{
		"title":   title,
		"content": content,
		"expires": expires,
	}

	var id int
	err := m.DB.QueryRow(context.Background(), stmt, args).Scan(&id)
	if err != nil {
		fmt.Println("Hello")
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
