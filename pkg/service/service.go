package service

import (
	todo "todo-app/app-models"
	"todo-app/clients/sso/grpc"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password_hash string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(userId int, userList todo.ToDoList) (int, error)
	GetAll(userId int) ([]todo.ToDoList, error)
	GetById(userId, listId int) (todo.ToDoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, updateData todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(userId, listId int, item todo.ToDoItem) (int, error)
	GetAll(userId, listId int) ([]todo.ToDoItem, error)
	GetById(userId, itemId int) (todo.ToDoItem, error)
	Delete(userId, listId int) error
	Update(userId, listId int, updateData todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository, ssoclient *grpc.Client) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, ssoclient),
		ToDoList:      NewToDoListService(repos.ToDoList),
		ToDoItem:      NewToDoItemService(repos.ToDoItem, repos.ToDoList),
	}
}
