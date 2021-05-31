package repository

import (
	"strconv"

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

func (r *CardPostgres) Create(input plantapi.Card) (int, error) {

	var id int

	query := `INSERT INTO public."Cards"(
		c_id, power, pr_v, sec_v, shield, date, author, bid, tested, pic, dir, info, conn_type, coil_num, type, file) VALUES (default, 
			  $1,    $2,   $3,    $4,     $5,   $6,     $7,  $8,     $9,  $10,  $11, $12,       $13,       $14, $15) returning c_id`

	row := r.db.QueryRow(query, input.Power, input.Pr_v, input.Sec_v, input.Shield, input.Date, input.Author, input.Bid, input.Tested, input.Pic, input.Dir, input.Info, input.Conn_type, input.Coil_num, input.Type, input.File)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	query = `UPDATE public."Dirs"
	SET last=$1
	WHERE d_id = $2 returning d_id`
	a, err := strconv.Atoi(input.Last)
	if err != nil {
		return 0, err
	}
	row = r.db.QueryRow(query, a+1, input.D_id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
