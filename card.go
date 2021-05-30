package plantapi

type Card struct {
	C_id      int     `json:"c_id" db:"c_id"`
	Type      string  `json:"type" db:"type"`
	Power     float32 `json:"power" db:"power"`
	Info      string  `json:"info" db:"info"`
	Pr_v      string  `json:"pr_v" db:"pr_v"`
	Sec_v     string  `json:"sec_v" db:"sec_v"`
	Shield    bool    `json:"shield" db:"shield"`
	Date      string  `json:"date" db:"date"`
	Author    string  `json:"author" db:"author"`
	Bid       bool    `json:"Bid" db:"Bid"`
	Tested    bool    `json:"tested" db:"tested"`
	Pic       string  `json:"pic" db:"pic"`
	Dir       string  `json:"dir" db:"dir"`
	Conn_type string  `json:"conn_type" db:"conn_type"`
	Coil_num  int     `json:"coil_num" db:"coil_num"`
	Mes       string  `json:"mes" db:"mes"`
}
