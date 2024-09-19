package routes

import (
	"chat_websocket/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	//Enrutamiento de Ejemplo
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	//Auth Routes
	auth := router.Group("/auth")
	{
		auth.GET("/login", controllers.ThisAuthController.Login)
		auth.POST("/register")
		auth.POST("/logout")
	}

}
