package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go-backend-task/models"
)

var users []models.User
var idCounter = 1

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// POST /users
func CreateUser(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
		DOB  string `json:"dob"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dob, err := time.Parse("2006-01-02", input.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB format"})
		return
	}

	user := models.User{
		ID:   idCounter,
		Name: input.Name,
		DOB:  dob,
	}

	idCounter++
	users = append(users, user)

	c.JSON(http.StatusCreated, gin.H{
		"id":   user.ID,
		"name": user.Name,
		"dob":  input.DOB,
	})
}

// GET /users/:id
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"id":   user.ID,
				"name": user.Name,
				"dob":  user.DOB.Format("2006-01-02"),
				"age":  calculateAge(user.DOB),
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input struct {
		Name string `json:"name"`
		DOB  string `json:"dob"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dob, err := time.Parse("2006-01-02", input.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB format"})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Name = input.Name
			users[i].DOB = dob

			c.JSON(http.StatusOK, gin.H{
				"id":   user.ID,
				"name": input.Name,
				"dob":  input.DOB,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// GET /users
func ListUsers(c *gin.Context) {
	var result []gin.H

	for _, user := range users {
		result = append(result, gin.H{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.DOB.Format("2006-01-02"),
			"age":  calculateAge(user.DOB),
		})
	}

	c.JSON(http.StatusOK, result)
}
