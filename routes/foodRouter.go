package routes

import (
	controller "github.com/danh996/Golang-Restaurant-Management-Backend-Project/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/food", controller.GetFoods())
	incomingRouter.GET("/foods/:food_id", controller.GetFood())
	incomingRouter.POST("/foods/:food_id", controller.CreateFood())
	incomingRouter.PATCH("/foods/:food_id", controller.UpdateFood())
}
