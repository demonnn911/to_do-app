package repository

import (
	"errors"
	"fmt"
	todo "todo-app/app-models"

	"github.com/jmoiron/sqlx"
)

type AuthSQL struct {
	db *sqlx.DB
}

func NewAuthSQL(db *sqlx.DB) *AuthSQL {
	return &AuthSQL{db: db}
}

func (r *AuthSQL) CreateUser(id int64) error {
	query := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", usersTable)
	row, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	res, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return errors.New("couldn't input user's data")
	}
	return nil
}

func (r *AuthSQL) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
