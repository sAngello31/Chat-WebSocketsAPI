package controllers

import (
	"chat_websocket/models"
	"chat_websocket/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserRepo *models.UserRepository
}

var ThisAuthController *AuthController

func NewAuthController(userRepo *models.UserRepository) {
	ThisAuthController = &AuthController{UserRepo: userRepo}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	//ctrl.UserRepo.Login(c)
}

func RegisterNewUser(c *gin.Context) {
	services.RegisterUser(c)
}
