package service

import (
	"context"
	"errors"
	todo "todo-app/app-models"
	"todo-app/pkg/repository"
)

type ToDoItemService struct {
	Repo     repository.ToDoItem
	ListRepo repository.ToDoList
}

func NewToDoItemService(repo repository.ToDoItem, listRepo repository.ToDoList) *ToDoItemService {
	return &ToDoItemService{
		Repo:     repo,
		ListRepo: listRepo,
	}
}

func (s *ToDoItemService) Create(ctx context.Context, userId, listId int64, item todo.ToDoItem) (int64, error) {
	_, err := s.ListRepo.GetById(ctx, userId, listId)
	if err != nil {
		return 0, errors.New("there is no list with such id, or user doesn't have permission for it")
	}
	return s.Repo.Create(ctx, listId, item)
}

func (s *ToDoItemService) GetAll(ctx context.Context, userId, listId int64) ([]todo.ToDoItem, error) {
	_, err := s.ListRepo.GetById(ctx, userId, listId)
	if err != nil {
		return nil, errors.New("there is no list with such id, or user doesn't have permission for it")
	}
	//	return s.Repo.GetAll(userId, listId)
	return s.Repo.GetAll(ctx, listId)

}

func (s *ToDoItemService) GetById(ctx context.Context, userId, itemId int64) (todo.ToDoItem, error) {
	return s.Repo.GetById(ctx, userId, itemId)
}

func (s *ToDoItemService) Delete(ctx context.Context, userId, itemId int64) error {
	return s.Repo.Delete(ctx, userId, itemId)
}

func (s *ToDoItemService) Update(ctx context.Context, userId, listId int64, updateData todo.UpdateItemInput) error {
	if err := updateData.Validate(); err != nil {
		return err
	}
	return s.Repo.Update(ctx, userId, listId, updateData)
}
