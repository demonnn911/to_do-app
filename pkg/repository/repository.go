package repository

import (
	todo "todo-app/app-models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type ToDoList interface {
	Create(userId int, list todo.ToDoList) (int, error)
	GetAll(userId int) ([]todo.ToDoList, error)
	GetById(userId, listId int) (todo.ToDoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, updateData todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(listId int, item todo.ToDoItem) (int, error)
	GetAll(listId int) ([]todo.ToDoItem, error)
	GetById(userId, itemId int) (todo.ToDoItem, error)
	Delete(userId, itemId int) error
	Update(userId, listId int, updateData todo.UpdateItemInput) error
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
