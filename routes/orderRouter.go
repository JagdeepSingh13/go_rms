package routes

import (
	"github.com/JagdeepSingh13/go_rms/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(routes *gin.Engine) {
	routes.GET("/orders", controllers.GetOrders())
	routes.GET("/orders/:order_id", controllers.GetOrder())
	routes.POST("/orders", controllers.CreateOrder())
	routes.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
