package routes

import (
	controller "github.com/danh996/Golang-Restaurant-Management-Backend-Project/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/invoices", controller.GetInvoices())
	incomingRouter.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRouter.POST("/invoices/:invoice_id", controller.CreateInvoice())
	incomingRouter.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
