package service

import (
	todo "todo-app/app-models"
	"todo-app/clients/sso/grpc"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int64, error)
	ValidateToken(token string) (int64, error)
	Login(input todo.SignInInput) (string, error)
}

type ToDoList interface {
	Create(userId int64, userList todo.ToDoList) (int64, error)
	GetAll(userId int64) ([]todo.ToDoList, error)
	GetById(userId, listId int64) (todo.ToDoList, error)
	Delete(userId, listId int64) error
	Update(userId, listId int64, updateData todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(userId, listId int64, item todo.ToDoItem) (int64, error)
	GetAll(userId, listId int64) ([]todo.ToDoItem, error)
	GetById(userId, itemId int64) (todo.ToDoItem, error)
	Delete(userId, listId int64) error
	Update(userId, listId int64, updateData todo.UpdateItemInput) error
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
