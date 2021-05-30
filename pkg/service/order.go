package service

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/daria40tim/plant-api/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetAll() ([]plantapi.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) GetById(c_id int) (plantapi.Order, error) {
	return s.repo.GetById(c_id)
}

func (s *OrderService) UpdateById(c_id int) (int, error) {
	return s.repo.UpdateById(c_id)
}

func (s *OrderService) Create(plantapi.Order) (int, error) {
	return s.repo.Create(plantapi.Order)
}
