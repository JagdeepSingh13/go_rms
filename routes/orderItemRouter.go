package routes

import (
	"github.com/JagdeepSingh13/go_rms/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(routes *gin.Engine) {
	routes.GET("/orderItems", controllers.GetOrderItems())
	routes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	routes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	routes.POST("/orderItems", controllers.CreateOrderItem())
	routes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
