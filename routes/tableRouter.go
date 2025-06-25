package routes

import (
	"github.com/JagdeepSingh13/go_rms/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(routes *gin.Engine) {
	routes.GET("/tables", controllers.GetTables())
	routes.GET("/tables/:table_id", controllers.GetTable())
	routes.POST("/tables", controllers.CreateTable())
	routes.PATCH("/tables/:table_id", controllers.UpdateTable())
}
