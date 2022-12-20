package routes

import "github.com/gin-gonic/gin"

func FoodRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
