package routes

import (
	controller "github.com/danh996/Golang-Restaurant-Management-Backend-Project/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/orders", controller.GetOrders())
	incomingRouter.GET("/orders/:order_id", controller.GetOrder())
	incomingRouter.POST("/orders/:order_id", controller.CreateOrder())
	incomingRouter.PATCH("/orders/:order_id", controller.UpdateOrder())
}
