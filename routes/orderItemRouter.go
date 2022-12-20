package routes

import (
	controller "github.com/danh996/Golang-Restaurant-Management-Backend-Project/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/orderItems", controller.GetOrderItems())
	incomingRouter.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRouter.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRouter.POST("/orderItems/:orderItem_id", controller.CreateOrderItem())
	incomingRouter.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
