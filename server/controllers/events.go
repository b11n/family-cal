package controllers

import (
	"net/http"
	"web-service/db"
	"web-service/models"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	var eventList []db.Event
	db.DB.Find(&eventList)
	c.IndentedJSON(http.StatusOK, eventList)
}

func CreateEvent(c *gin.Context) {
	var newEvent models.Event

	err := c.BindJSON(&newEvent)

	if err != nil {
		return
	}

	newEventForDb := db.Event{Title: newEvent.Title, From: newEvent.From, To: newEvent.To, CalendarID: "bn"}
	db.DB.Create(&newEventForDb)

	c.IndentedJSON(http.StatusCreated, newEventForDb)
}

func GetEventByID(c *gin.Context) {
	id := c.Param("id")

	var event db.Event

	err := db.DB.Where("id = ?", id).First(&event).Error

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, event)
}
