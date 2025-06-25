package routes

import (
	"github.com/JagdeepSingh13/go_rms/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.GET("/users", controllers.GetUsers())
	routes.GET("/users/:user_id", controllers.GetUser())
	routes.POST("/users/signup", controllers.SignUp())
	routes.GET("/users/login", controllers.Login())
}
