package controllers

import (
	"chat_websocket/models"

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
	c.JSON(200, gin.H{
		"message": "Te voy a loggear bitch",
	})
}
