package routes

import (
	"github.com/JagdeepSingh13/go_rms/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(routes *gin.Engine) {
	routes.GET("/foods", controllers.GetFoods())
	routes.GET("/foods/:food_id", controllers.GetFood())
	routes.POST("/foods", controllers.CreateFood())
	routes.PATCH("/foods/:food_id", controllers.UpdateFood())
}
