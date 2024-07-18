package models

import (
	"database/sql"
	"github.com/https_whoyan/Lets_Go_Book_Course/pkg/repository/postgres"
	"time"
)

type Snippet struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type SnippetModel struct {
	db *sql.DB
}

func NewSnippetModel(serverAddr string) (*SnippetModel, error) {
	db, err := postgres.Open(serverAddr)
	if err != nil {
		return nil, nil
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &SnippetModel{db: db}, nil
}

// Insert TODO
func (m *SnippetModel) Insert(title string, content string, expiresAt int) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return []*Snippet{}, nil
}
