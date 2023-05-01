package models

import "time"

type Event struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
}
