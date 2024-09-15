package users

import "database/sql"

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
	return nil
}

func (m *UsersModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UsersModel) Exists(id int) (bool, error) {
	return false, nil
}
