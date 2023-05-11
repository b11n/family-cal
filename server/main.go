package main

import (
	"web-service/controllers"
	"web-service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.ConnectDatabase()

	router.GET("/events", controllers.GetEvents)
	router.POST("/event", controllers.CreateEvent)
	router.GET("/event/:id", controllers.GetEventByID)
	router.GET("events/calendar/:id", controllers.GetEventsByCalendarID)

	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:id", controllers.GetUser)

	router.POST("/calendar", controllers.CreateCalendar)
	router.GET("/calendar/:id", controllers.GetCalendar)

	router.Run("0.0.0.0:8080")
}
