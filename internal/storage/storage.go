package storage

import (
	model "github.com/DeMarDeXis/VProj/internal/model"
	"github.com/DeMarDeXis/VProj/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type NHLList interface {
	GetTeams() ([]model.NHLTeamsOutput, error)
	GetSchedule() ([]model.NHLScheduleOutput, error)
	GetLastSchedule(count int) ([]model.NHLScheduleOutput, error)
}

type NBAList interface {
	// GetTeams() (error) //TODO: add model
}

type Storage struct {
	Authorization
	NHLList
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: postgres.NewAuth(db),
		NHLList:       postgres.NewNHLList(db),
	}
}
