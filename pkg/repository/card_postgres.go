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

func (r *CardPostgres) GetAll() ([]plantapi.CardAll, error) {
	var res []plantapi.CardAll

	query := `SELECT c_id, power, pr_v, sec_v, shield, date, bid, tested, dir, info, conn_type, coil_num, name as type, file, case mes when true then 'кВа' else 'кВт' end as mes
	FROM public."Cards"
	left join public."Types" on type = t_id`

	err := r.db.Select(&res, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *CardPostgres) GetById(c_id int) (plantapi.Card, error) {
	var res plantapi.Card
	return res, nil
}

func (r *CardPostgres) UpdateById(c_id int) (int, error) {

	return c_id, nil
}

func (r *CardPostgres) Create(input plantapi.Card) (int, error) {

	var c_id int
	var id int

	query := `INSERT INTO public."Cards"(
		c_id, power, pr_v, sec_v, shield, date, author, bid, tested,  dir, info, conn_type, coil_num, type, file, mes) VALUES (default, 
			  $1,    $2,   $3,    $4,     $5,   $6,     $7,  $8,      $9,  $10, $11,       $12,       $13, $14, $15) returning c_id`

	row := r.db.QueryRow(query, input.Power, input.Pr_v, input.Sec_v, input.Shield, input.Date, input.Author, input.Bid, input.Tested, input.Dir, input.Info, input.Conn_type, input.Coil_num, input.Type, input.File, input.Mes)
	if err := row.Scan(&c_id); err != nil {
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

	query = `INSERT INTO public."Pics"(
		pic_id, pic, c_id)
		VALUES (default, $1, $2) returning pic_id`

	row = r.db.QueryRow(query, input.Pic, c_id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return c_id, nil
}
