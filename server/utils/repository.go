package utils

import (
	"chat_websocket/controllers"
	"chat_websocket/models"
)

func InitRepositories() {
	//Repositories
	userRepo := models.NewUserRepository(Client)
	chatGroupRepo := models.NewChatGroupRepository(Client)
	MsgRepo := models.NewMsgRepository(Client)
	//Controllers
	controllers.NewAuthController(userRepo)
	controllers.NewChatGroupController(chatGroupRepo)
	controllers.NewChatController(MsgRepo)
}
