package repository

import (
	todo "todo-app/app-models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(id int64) error
	GetUser(username, password string) (todo.User, error)
}

type ToDoList interface {
	Create(userId int64, list todo.ToDoList) (int64, error)
	GetAll(userId int64) ([]todo.ToDoList, error)
	GetById(userId, listId int64) (todo.ToDoList, error)
	Delete(userId, listId int64) error
	Update(userId, listId int64, updateData todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(listId int64, item todo.ToDoItem) (int64, error)
	GetAll(listId int64) ([]todo.ToDoItem, error)
	GetById(userId, itemId int64) (todo.ToDoItem, error)
	Delete(userId, itemId int64) error
	Update(userId, listId int64, updateData todo.UpdateItemInput) error
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
