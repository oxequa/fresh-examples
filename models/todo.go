package models

type Todo struct {
	Uuid   string `json:"uuid"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
