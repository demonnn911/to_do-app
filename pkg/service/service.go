package service

import (
	"context"
	todo "todo-app/app-models"
	"todo-app/clients/sso/grpc"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, user todo.User) (int64, error)
	ValidateToken(ctx context.Context, token string) (int64, error)
	Login(ctx context.Context, input todo.SignInInput) (string, error)
}

type ToDoList interface {
	Create(ctx context.Context, userId int64, userList todo.ToDoList) (int64, error)
	GetAll(ctx context.Context, userId int64) ([]todo.ToDoList, error)
	GetById(ctx context.Context, userId, listId int64) (todo.ToDoList, error)
	Delete(ctx context.Context, userId, listId int64) error
	Update(ctx context.Context, userId, listId int64, updateData todo.UpdateListInput) error
}

type ToDoItem interface {
	Create(ctx context.Context, userId, listId int64, item todo.ToDoItem) (int64, error)
	GetAll(ctx context.Context, userId, listId int64) ([]todo.ToDoItem, error)
	GetById(ctx context.Context, userId, itemId int64) (todo.ToDoItem, error)
	Delete(ctx context.Context, userId, listId int64) error
	Update(ctx context.Context, userId, listId int64, updateData todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository, client *grpc.SSOClientWrapper) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, client.SSOProvider),
		ToDoList:      NewToDoListService(repos.ToDoList),
		ToDoItem:      NewToDoItemService(repos.ToDoItem, repos.ToDoList),
	}
}
