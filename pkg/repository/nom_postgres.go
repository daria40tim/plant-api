package repository

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/jmoiron/sqlx"
)

type NomPostgres struct {
	db *sqlx.DB
}

func NewNomPostgres(db *sqlx.DB) *NomPostgres {
	return &NomPostgres{db: db}
}

func (r *NomPostgres) GetAll() ([]plantapi.Nom, error) {

	return nil, nil
}

func (r *NomPostgres) GetById(c_id int) (plantapi.Nom, error) {
	var res plantapi.Nom
	return res, nil
}

func (r *NomPostgres) UpdateById(c_id int) (int, error) {

	return 0, nil
}

func (r *NomPostgres) Create(plantapi.Nom) (int, error) {

	return 0, nil
}
