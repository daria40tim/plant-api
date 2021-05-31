package service

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/daria40tim/plant-api/pkg/repository"
)

type CardService struct {
	repo repository.Card
}

func NewCardService(repo repository.Card) *CardService {
	return &CardService{repo: repo}
}

func (s *CardService) GetAll() ([]plantapi.CardAll, error) {
	return s.repo.GetAll()
}

func (s *CardService) GetById(c_id int) (plantapi.Card, error) {
	return s.repo.GetById(c_id)
}

func (s *CardService) UpdateById(c_id int) (int, error) {
	return s.repo.UpdateById(c_id)
}

func (s *CardService) Create(card plantapi.Card) (int, error) {
	return s.repo.Create(card)
}
