package plantapi

type Nom struct {
	N_id    int    `json:"n_id" db:"n_id"`
	Title   string `json:"title" db:"title"`
	Feature string `json:"feature" db:"feature"`
	Mes     string `json:"mes" db:"mes"`
}
