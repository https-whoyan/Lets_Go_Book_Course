package users

import (
	"database/sql"
	"errors"
	myErrors "github.com/https_whoyan/Lets_Go_Book_Course/internal/errors"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/models"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/usecases/hash"
)

type UsersModel struct {
	db *sql.DB
}

func NewUsersModel(db *sql.DB) (*UsersModel, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return &UsersModel{db: db}, nil
}

func (m *UsersModel) Insert(name string, mail string, pass string) error {
	hashedPassword := hash.GetHash(pass)
	_, err := m.db.Exec(insertUserStatement, name, mail, hashedPassword)
	if err != nil {
		if myErrors.CheckIsSQLUniqueConstrainsError(err) {
			return myErrors.ErrDuplicateEmail
		}
		return err
	}
	return nil
}

func (m *UsersModel) Authenticate(email, password string) (int, error) {
	var user = &models.User{}
	err := m.db.QueryRow(selectByEmail, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	passwordHash := hash.GetHash(password)
	if passwordHash == string(user.HashedPassword) {
		return user.ID, nil
	}
	return 0, myErrors.ErrInvalidCredentials
}

func (m *UsersModel) Exists(id int) (bool, error) {
	rows, err := m.db.Query(selectById, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	defer rows.Close()
	return true, nil
}
