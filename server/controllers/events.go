package controllers

import (
	"net/http"
	"web-service/db"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	var eventList []db.Event
	db.DB.Find(&eventList)
	c.IndentedJSON(http.StatusOK, eventList)
}

func CreateEvent(c *gin.Context) {
	var newEvent db.Event

	err := c.BindJSON(&newEvent)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	db.DB.Create(&newEvent)

	c.IndentedJSON(http.StatusCreated, newEvent)
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

func GetEventsByCalendarID(c *gin.Context) {
	id := c.Param("id")

	var calendar []db.Event

	err := db.DB.Where("calendar_id = ?", id).Find(&calendar).Error

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "calendar not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, calendar)
}
