package service

import (
	"errors"
	"todo-app"
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

func (s *ToDoItemService) Create(userId, listId int, item todo.ToDoItem) (int, error) {
	_, err := s.ListRepo.GetById(userId, listId)
	if err != nil {
		return 0, errors.New("there is no list with such id, or user doesn't have permission for it")
	}
	return s.Repo.Create(listId, item)
}

func (s *ToDoItemService) GetAll(userId, listId int) ([]todo.ToDoItem, error) {
	_, err := s.ListRepo.GetById(userId, listId)
	if err != nil {
		return nil, errors.New("there is no list with such id, or user doesn't have permission for it")
	}
	//	return s.Repo.GetAll(userId, listId)
	return s.Repo.GetAll(listId)

}

func (s *ToDoItemService) GetById(userId, itemId int) (todo.ToDoItem, error) {
	return s.Repo.GetById(userId, itemId)
}

func (s *ToDoItemService) Delete(userId, itemId int) error {
	return s.Repo.Delete(userId, itemId)
}

func (s *ToDoItemService) Update(userId, listId int, updateData todo.UpdateItemInput) error {
	if err := updateData.Validate(); err != nil {
		return err
	}
	return s.Repo.Update(userId, listId, updateData)
}
