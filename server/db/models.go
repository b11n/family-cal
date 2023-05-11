package db

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type Event struct {
	BaseModel
	Title      string    `json:"title"`
	From       time.Time `json:"from"`
	To         time.Time `json:"to"`
	CalendarID int       `json:"calendarId"`
	Calendar   Calendar  `json:"-"`
}

type Calendar struct {
	BaseModel
	Name   string `json:"name"`
	UserID int    `json:"userId"`
	User   User   `json:"-"`
}

type User struct {
	BaseModel
	Name  string `json:"name"`
	Email string `json:"email"`
}
