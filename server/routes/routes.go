package routes

import (
	"chat_websocket/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.RegisterNewUser)
		auth.POST("/logout")
	}

	user := router.Group("/user")
	{
		user.GET("/getAll", controllers.GetAllUsers)
	}

	chat := router.Group("/chat")
	{
		chat.GET("/getUUID/:userA/:userB", controllers.GetUUIDUsers)
		chat.GET("/connectTo/:uuid", controllers.HandleConn)
	}

}
