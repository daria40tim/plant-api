package repository

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/jmoiron/sqlx"
)

type Card interface {
	GetAll() ([]plantapi.Card, error)
	GetById(c_id int) (plantapi.Card, error)
	UpdateById(c_id int) (int, error)
	Create(plantapi.Card) (int, error)
}

type Nom interface {
	GetAll() ([]plantapi.Nom, error)
	GetById(c_id int) (plantapi.Nom, error)
	UpdateById(c_id int) (int, error)
	Create(plantapi.Nom) (int, error)
}

type Order interface {
	GetAll() ([]plantapi.Order, error)
	GetById(c_id int) (plantapi.Order, error)
	UpdateById(c_id int) (int, error)
	Create(plantapi.Order) (int, error)
}

type OCR interface {
	GetDirs() ([]plantapi.Dir, error)
	GetImg(index int, dir string) (plantapi.CardOCR, error)
}

type Repository struct {
	Card
	Nom
	Order
	OCR
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Card:  NewCardPostgres(db),
		Nom:   NewNomPostgres(db),
		Order: NewOrderPostgres(db),
		OCR:   NewOCRPostgres(db),
	}
}
