package repository

import (
	"context"
	"time"
	"todo-app/pkg/config/env"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListTable   = "todo_list"
	usersListsTable = "users_lists"
	todoItemTable   = "todo_item"
	listItemsTable  = "list_items"
)

func NewPostgresDB(cfg env.DBConfig) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := sqlx.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
