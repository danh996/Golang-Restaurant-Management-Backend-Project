package routes

import (
	"github.com/danh996/Golang-Restaurant-Management-Backend-Project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRouter *gin.Engine) {
	incomingRouter.GET("/users", controllers.GetUsers())
	incomingRouter.GET("/users/:user_id", controllers.GetUser())
	incomingRouter.POST("/users/signup", controllers.SignUp())
	incomingRouter.POST("/users/login", controllers.Login())
}
