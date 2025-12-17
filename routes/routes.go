package routes

import (
	"github.com/gin-gonic/gin"
	"go-backend-task/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
}
