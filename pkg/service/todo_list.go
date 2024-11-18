package service

import (
	"context"
	todo "todo-app/app-models"
	"todo-app/pkg/repository"
)

type ToDoListService struct {
	repo repository.ToDoList
}

func NewToDoListService(repo repository.ToDoList) *ToDoListService {
	return &ToDoListService{
		repo: repo,
	}
}

func (s *ToDoListService) Create(ctx context.Context, userId int64, list todo.ToDoList) (int64, error) {
	return s.repo.Create(ctx, userId, list)
}

func (s *ToDoListService) GetAll(ctx context.Context, userId int64) ([]todo.ToDoList, error) {
	return s.repo.GetAll(ctx, userId)
}

func (s *ToDoListService) GetById(ctx context.Context, userId, listId int64) (todo.ToDoList, error) {
	return s.repo.GetById(ctx, userId, listId)
}

func (s *ToDoListService) Delete(ctx context.Context, userId, listId int64) error {
	return s.repo.Delete(ctx, userId, listId)
}

func (s *ToDoListService) Update(ctx context.Context, userId, listId int64, updateData todo.UpdateListInput) error {
	if err := updateData.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, userId, listId, updateData)
}
