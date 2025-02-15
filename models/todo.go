package models

type ToDo struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
