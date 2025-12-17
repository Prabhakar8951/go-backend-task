package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-backend-task/models"
)

var users []models.User

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}
