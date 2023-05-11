package models

import "time"

type Event struct {
	ID    		string    `json:"id"`
	Title 		string    `json:"title"`
	From  		time.Time `json:"from"`
	To    		time.Time `json:"to"`
	CalendarID	string 	  `json: "calendar"`
}

type Calendar struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
	UserID 	string `json:"userId"` 
}

type User struct {
	ID 		string `json:"id"`
	Name	string `json:"name"`
	Email   string `json:"email"`
}


