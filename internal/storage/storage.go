package storage

import (
	model "github.com/DeMarDeXis/VProj/internal/model"
	"github.com/DeMarDeXis/VProj/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username string) (model.User, error)
}

type TodoList interface {
	Create(userID int, list model.TodoList) (int, error)
	GetAll(userID int) ([]model.TodoList, error)
	GetByID(userID, listID int) (model.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input model.UpdateListInput) error
}

type Storage struct {
	Authorization
	TodoList
}

func NewStorage(db *sqlx.DB, log *slog.Logger) *Storage {
	return &Storage{
		Authorization: postgres.NewAuth(db),
		TodoList:      postgres.NewTodoList(db, log),
	}
}
