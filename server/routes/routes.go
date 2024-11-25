package routes

import (
	"chat_websocket/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	//Auth Routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.ThisAuthController.Login)
		auth.POST("/register", controllers.RegisterNewUser)
		auth.POST("/logout")
	}

	chatGroup := router.Group("/chatgroup")
	{
		chatGroup.GET("/getAllChatGroups", controllers.ThisChatGroupController.GetAllChatsGroups)
		chatGroup.POST("/insertChat", controllers.ThisChatGroupController.InsertChatGroup)
	}

	chat := router.Group("/chat")
	{
		chat.GET("connectTo", controllers.ThisChatController.InitChat)
	}

}
