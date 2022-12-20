package routes

import (
	controller "github.com/danh996/Golang-Restaurant-Management-Backend-Project/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/menus", controller.GetMenus())
	incomingRouter.GET("/menus/:menu_id", controller.GetMenu())
	incomingRouter.POST("/menus/:menu_id", controller.CreateMenu())
	incomingRouter.PATCH("/menus/:menu_id", controller.UpdateMenu())
}
