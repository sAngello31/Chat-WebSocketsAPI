package utils

import (
	"chat_websocket/controllers"
	"chat_websocket/models"
)

func InitRepositories() {
	//Repositories
	userRepo := models.NewUserRepository(Client)
	//Controllers
	controllers.NewAuthController(userRepo)
}
