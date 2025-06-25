package routes

import (
	"github.com/JagdeepSingh13/go_rms/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(routes *gin.Engine) {
	routes.GET("/menus", controllers.GetMenus())
	routes.GET("/menus/:menu_id", controllers.GetMenu())
	routes.POST("/menus", controllers.CreateMenu())
	routes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
