package service

import (
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/DeMarDeXis/VProj/internal/storage"
)

type NHLListService struct {
	storage storage.NHLList
}

func NewNHLListService(storage storage.NHLList) *NHLListService {
	return &NHLListService{storage: storage}
}

func (s *NHLListService) GetTeams() ([]model.NHLTeamsOutput, error) {
	return s.storage.GetTeams()
}

func (s *NHLListService) GetSchedule() ([]model.NHLScheduleOutput, error) {
	return s.storage.GetSchedule()
}

func (s *NHLListService) GetLastSchedule(count int) ([]model.NHLScheduleOutput, error) {
	return s.storage.GetLastSchedule(count)
}
