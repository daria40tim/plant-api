package repository

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/jmoiron/sqlx"
)

type Cards interface {
	GetAll() ([]plantapi.Card, error)
	GetById(c_id int) (plantapi.Card, error)
	UpdateById(c_id int) (int, error)
	Create(plantapi.Card) (int, error)
}

type Noms interface {
	GetAll() ([]plantapi.Nom, error)
	GetById(c_id int) (plantapi.Nom, error)
	UpdateById(c_id int) (int, error)
	Create(plantapi.Nom) (int, error)
}

type Orders interface {
	GetAll() ([]plantapi.Order, error)
	GetById(c_id int) (plantapi.Order, error)
	UpdateById(c_id int) (int, error)
	Create(plantapi.Order) (int, error)
}

type Repository struct {
	Cards
	Noms
	Orders
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Cards:  NewCardPostgres(db),
		Noms:   NewNomPostgres(db),
		Orders: NewOrderPostgres(db),
	}
}
