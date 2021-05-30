package repository

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/jmoiron/sqlx"
)

type CardPostgres struct {
	db *sqlx.DB
}

func NewCardPostgres(db *sqlx.DB) *CardPostgres {
	return &CardPostgres{db: db}
}

func (r *CardPostgres) GetAll() ([]plantapi.Card, error) {

	return nil, nil
}

func (r *CardPostgres) GetById(c_id int) (plantapi.Card, error) {
	var res plantapi.Card
	return res, nil
}

func (r *CardPostgres) UpdateById(c_id int) (int, error) {

	return c_id, nil
}

func (r *CardPostgres) Create(plantapi.Card) (int, error) {

	return 0, nil
}
