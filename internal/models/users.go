package models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	CreatedAt      time.Time
}

type UsersModel struct {
	db *sql.DB
}

func (m *UsersModel) Insert(name string, mail string, pass string) (*User, error) {}
