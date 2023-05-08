package controllers

import (
	"net/http"
	"time"
	"web-service/models"

	"github.com/gin-gonic/gin"
)

var events = []models.Event{
	{ID: "1", Title: "Stand up", From: time.Now(), To: time.Now()},
	{ID: "2", Title: "Gym", From: time.Now(), To: time.Now()},
	{ID: "4", Title: "Meeting", From: time.Now(), To: time.Now()},
}

func GetEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func CreateEvent(c *gin.Context) {
	var newEvent models.Event

	err := c.BindJSON(&newEvent)

	if err != nil {
		return
	}

	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func GetEventByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range events {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
