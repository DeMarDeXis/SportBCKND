package service

import (
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/DeMarDeXis/VProj/internal/storage"
)

type TodoListService struct {
	storage storage.TodoList
}

func NewTodoListService(storage storage.TodoList) *TodoListService {
	return &TodoListService{storage: storage}
}

func (s *TodoListService) Create(userID int, list model.TodoList) (int, error) {
	return s.storage.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]model.TodoList, error) {
	return s.storage.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (model.TodoList, error) {
	return s.storage.GetByID(userID, listID)
}

func (s *TodoListService) Delete(userID, listID int) error {
	return s.storage.Delete(userID, listID)
}

func (s *TodoListService) Update(userID, listID int, input model.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.storage.Update(userID, listID, input)
}
