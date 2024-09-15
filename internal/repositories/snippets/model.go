package snippets

import (
	"database/sql"
	"errors"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/models"

	myErrors "github.com/https_whoyan/Lets_Go_Book_Course/internal/errors"
)

type SnippetModel struct {
	db *sql.DB
}

func NewSnippetModel(db *sql.DB) (*SnippetModel, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return &SnippetModel{db: db}, nil
}

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

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	s := &models.Snippet{}

	err := m.db.QueryRow(selectSnippetStatement, id).
		Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, myErrors.ErrNoRecords
		}
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	var snippets []*models.Snippet

	rows, err := m.db.Query(multipleSelectSnippet)
	if err != nil {
		return []*models.Snippet{}, err
	}
	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		s := &models.Snippet{}
		internalErr := rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
		if internalErr != nil {
			return []*models.Snippet{}, internalErr
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return []*models.Snippet{}, err
	}
	return snippets, nil
}
