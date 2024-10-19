package service

import (
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/DeMarDeXis/VProj/internal/storage"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list model.TodoList) (int, error)
	GetAll(userID int) ([]model.TodoList, error)
	GetByID(userID, listID int) (model.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input model.UpdateListInput) error
}

type Service struct {
	Auth
	TodoList
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Auth:     NewAuthService(storage.Authorization),
		TodoList: NewTodoListService(storage.TodoList),
	}
}
