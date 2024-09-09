package service

import (
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/DeMarDeXis/VProj/internal/storage"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username string) (string, error)
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

func NewService(strg *storage.Storage) *Service {
	return &Service{
		Auth:     NewAuthService(strg.Authorization),
		TodoList: NewTodoListService(strg.TodoList),
	}
}
