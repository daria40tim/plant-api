package plantapi

type Dir struct {
	D_id     int    `json:"d_id" db:"d_id"`
	Name     string `json:"name" db:"name"`
	Count    int    `json:"count" db:"count"`
	Ch_count int    `json:"ch_count" db:"ch_count"`
	Last     int    `json:"last" db:"last"`
}

type CardOCR struct {
	Type  string `json:"type" db:"type"`
	Power string `json:"power" db:"power"`
	Pr_v  string `json:"pr_v" db:"pr_v"`
	Sec_v string `json:"sec_v" db:"sec_v"`
	Date  string `json:"date" db:"date"`
	Mes   string `json:"mes" db:"mes"`
	Text  string `json:"text" db:"text"`
	Img   string `json:"img" db:"img"`
	Name  string `json:"name" db:"name"`
}
