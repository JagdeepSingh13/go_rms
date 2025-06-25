package routes

import (
	"github.com/JagdeepSingh13/go_rms/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(routes *gin.Engine) {
	routes.GET("/invoices", controllers.GetInvoices())
	routes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	routes.POST("/invoices", controllers.CreateInvoice())
	routes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
