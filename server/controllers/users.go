package controllers

import (
	"net/http"
	"web-service/db"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser db.User

	err := c.BindJSON(&newUser)

	if err != nil {
		return
	}

	db.DB.Create(&newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user db.User

	err := db.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
