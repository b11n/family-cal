package main

import (
	"web-service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/events", controllers.GetEvents)
	router.POST("/event", controllers.CreateEvent)
	router.GET("/event/:id", controllers.GetEventByID)

	router.Run("localhost:8080")
}
