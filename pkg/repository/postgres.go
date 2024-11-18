package repository

import (
	"fmt"
	"todo-app/pkg/config"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListTable   = "todo_list"
	usersListsTable = "users_lists"
	todoItemTable   = "todo_item"
	listItemsTable  = "list_items"
)

func NewPostgresDB(cfg *config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port,
		cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
