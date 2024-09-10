package service

import (
	"encoding/json"
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/DeMarDeXis/VProj/internal/storage"
	"time"
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
	lists, err := s.storage.GetAll(userID)
	if err != nil {
		return nil, err
	}

	for i := range lists {
		dueDate, err := time.Parse(time.RFC3339, string(lists[i].DoeDate))
		if err != nil {
			return nil, err
		}
		dateInput := model.DateInput{
			Year:  dueDate.Year(),
			Month: int(dueDate.Month()),
			Day:   dueDate.Day(),
		}
		dateJSON, err := json.Marshal(dateInput)
		if err != nil {
			return nil, err
		}
		lists[i].DoeDate = json.RawMessage(dateJSON)
	}

	return lists, nil
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
