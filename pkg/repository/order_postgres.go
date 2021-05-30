package repository

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) GetAll() ([]plantapi.Order, error) {

	return nil, nil
}

func (r *OrderPostgres) GetById(c_id int) (plantapi.Order, error) {
	var res plantapi.Order
	return res, nil
}

func (r *OrderPostgres) UpdateById(c_id int) (int, error) {

	return 0, nil
}

func (r *OrderPostgres) Create(plantapi.Order) (int, error) {

	return 0, nil
}
