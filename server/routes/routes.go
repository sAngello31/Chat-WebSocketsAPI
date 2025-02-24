package routes

import (
	"chat_websocket/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/login", services.LoginUser)
		auth.POST("/register", services.RegisterUser)
		auth.POST("/logout")
	}

	user := router.Group("/user")
	{
		user.GET("/getAll", services.GetAllUsers)
		user.GET("/data", services.GetUserData)
	}

	chat := router.Group("/chat")
	{
		chat.GET("/connectTo/:uuid", services.HandleConn)
	}

}
