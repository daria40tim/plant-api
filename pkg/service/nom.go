package service

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/daria40tim/plant-api/repository"
)

type NomService struct {
	repo repository.Nom
}

func NewNomService(repo repository.Nom) *NomService {
	return &NomService{repo: repo}
}

func (s *NomService) GetAll() ([]plantapi.Nom, error) {
	return s.repo.GetAll()
}

func (s *NomService) GetById(c_id int) (plantapi.Nom, error) {
	return s.repo.GetById(c_id)
}

func (s *NomService) UpdateById(c_id int) (int, error) {
	return s.repo.UpdateById(c_id)
}

func (s *NomService) Create(plantapi.Nom) (int, error) {
	return s.repo.Create(plantapi.Nom)
}
