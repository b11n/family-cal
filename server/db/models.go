package db

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title      string
	From       time.Time
	To         time.Time
	CalendarID string
}
