package plantapi

type Card_nom struct {
	C_id    int     `json:"c_id" db:"c_id"`
	N_id    int     `json:"n_id" db:"n_id"`
	Title   string  `json:"title" db:"title"`
	Feature string  `json:"feature" db:"feature"`
	Mes     string  `json:"mes" db:"mes"`
	Count   float32 `json:"count" db:"count"`
}
