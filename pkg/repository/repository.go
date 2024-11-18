package repository

import (
	"context"
	todo "todo-app/app-models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, username, password string) (todo.User, error)
}

type ToDoList interface {
	Create(ctx context.Context, userId int64, list todo.ToDoList) (int64, error)
	GetAll(ctx context.Context, userId int64) ([]todo.ToDoList, error)
	GetById(ctx context.Context, userId, listId int64) (todo.ToDoList, error)
	Delete(ctx context.Context, userId, listId int64) error
	Update(ctx context.Context, userId, listId int64, updateData todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(ctx context.Context, listId int64, item todo.ToDoItem) (int64, error)
	GetAll(ctx context.Context, listId int64) ([]todo.ToDoItem, error)
	GetById(ctx context.Context, userId, itemId int64) (todo.ToDoItem, error)
	Delete(ctx context.Context, userId, itemId int64) error
	Update(ctx context.Context, userId, listId int64, updateData todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
		ToDoList:      NewToDoListPostgres(db),
		ToDoItem:      NewToDoItemPostgres(db),
	}
}
