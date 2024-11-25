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

func Login(c *gin.Context) {
	services.LoginUser(c)
}

func RegisterNewUser(c *gin.Context) {
	services.RegisterUser(c)
}
