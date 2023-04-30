package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type event struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	From  string `json:"from"`
	To    string `json:"to"`
}

var events = []event{
	{ID: "1", Title: "Stand up", From: "22-03-2022 11:15AM", To: "22-03-2022 11:15AM"},
	{ID: "2", Title: "Gym", From: "22-03-2022 11:15AM", To: "22-03-2022 11:15AM"},
}

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var newEvent event

	err := c.BindJSON(&newEvent)

	if err != nil {
		return
	}

	events = append(events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func getEventByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range events {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func main() {
	router := gin.Default()
	router.GET("/events", getEvents)
	router.POST("/event", createEvent)
	router.GET("/event/:id", getEventByID)

	router.Run("localhost:8080")
}
