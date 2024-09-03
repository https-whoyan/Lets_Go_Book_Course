package models

import (
	"database/sql"
	"errors"
	"fmt"
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

var (
	ErrNoRecords = fmt.Errorf("models: no matching record found")
)

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
	var id int
	row := m.db.QueryRow(insertSnippetStatement, title, content, expiresAt)
	if row.Err() != nil {
		return 0, row.Err()
	}
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	s := &Snippet{}

	err := m.db.QueryRow(selectSnippetStatement, id).
		Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecords
		}
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	var snippets []*Snippet

	rows, err := m.db.Query(multipleSelectSnippet)
	if err != nil {
		return []*Snippet{}, err
	}
	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		s := &Snippet{}
		internalErr := rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
		if internalErr != nil {
			return []*Snippet{}, internalErr
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return []*Snippet{}, err
	}
	return snippets, nil
}
