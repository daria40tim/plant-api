package service

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/daria40tim/plant-api/repository"
)

type Cards interface {
	GetAll() ([]plantapi.Card, error)
	GetById(c_id int) (plantapi.Card, error)
	UpdateById(c_id int) (int, error)
	AddById(plantapi.Card) (int, error)
}

type Noms interface {
	GetAll() ([]plantapi.Nom, error)
	GetById(c_id int) (plantapi.Nom, error)
	UpdateById(c_id int) (int, error)
	AddById(plantapi.Nom) (int, error)
}

type Orders interface {
	GetAll() ([]plantapi.Order, error)
	GetById(c_id int) (plantapi.Order, error)
	UpdateById(c_id int) (int, error)
	AddById(plantapi.Order) (int, error)
}

type Service struct {
	Cards
	Noms
	Orders
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Cards:  NewCardService(repos.Card),
		Noms:   NewNomService(repos.Nom),
		Orders: NewOrderService(repos.Order),
	}
}
