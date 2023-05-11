package controllers

import (
	"net/http"
	"web-service/db"

	"github.com/gin-gonic/gin"
)

func CreateCalendar(c *gin.Context) {
	var newCalendar db.Calendar

	err := c.BindJSON(&newCalendar)

	if err != nil {
		return
	}

	db.DB.Create(&newCalendar)

	c.IndentedJSON(http.StatusCreated, newCalendar)
}

func GetCalendar(c *gin.Context) {
	id := c.Param("id")

	var calendar db.Calendar

	err := db.DB.Where("id = ?", id).First(&calendar).Error

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "calendar not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, calendar)
}
