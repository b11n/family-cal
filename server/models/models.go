package models

type Event struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	From  string `json:"from"`
	To    string `json:"to"`
}
