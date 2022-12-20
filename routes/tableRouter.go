package routes

import (
	controller "github.com/danh996/Golang-Restaurant-Management-Backend-Project/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/tables", controller.GetTables())
	incomingRouter.GET("/tables/:table_id", controller.GetTable())
	incomingRouter.POST("/tables/:table_id", controller.CreateTable())
	incomingRouter.PATCH("/tables/:table_id", controller.UpdateTable())
}
