package repository

import (
	"encoding/base64"
	"io/ioutil"
	"log"

	plantapi "github.com/daria40tim/plant-api"
	"github.com/jmoiron/sqlx"
)

type OCRPostgres struct {
	db *sqlx.DB
}

func NewOCRPostgres(db *sqlx.DB) *OCRPostgres {
	return &OCRPostgres{db: db}
}

func (r *OCRPostgres) GetDirs() ([]plantapi.Dir, error) {
	var res []plantapi.Dir

	query := `SELECT d_id, count, ch_count, last, name
	FROM public."Dirs";`

	err := r.db.Select(&res, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *OCRPostgres) GetImg(index int, dir string) (plantapi.CardOCR, error) {
	var res plantapi.CardOCR
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadFile(dir + files[index].Name())
	if err != nil {
		log.Fatal(err)
	}

	base64Encoding := base64.StdEncoding.EncodeToString(bytes)

	res.Img = base64Encoding
	res.Name = files[index].Name()

	return res, nil
}
