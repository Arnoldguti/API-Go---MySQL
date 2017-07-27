package core



type Data struct {
	Id               int       `json:"id"`
	Description      string    `json:"description"`
	Username         string     `json:"username"`
	Element             string     `json:"element"`
	Date                string       `json:"date"`
}

type Datas []Data