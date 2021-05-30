package plantapi

type Order struct {
	O_id  int    `json:"o_id" db:"o_id"`
	Date  string `json:"date" db:"date"`
	Cards []Card `json:"cards"`
}
