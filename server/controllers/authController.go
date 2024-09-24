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
	ctrl.UserRepo.Login(c)
}

func (ctrl *AuthController) Register(c *gin.Context) {
	ctrl.UserRepo.RegisterNewUser(c)
}

func (ctrl *AuthController) GetUsersController(c *gin.Context) {
	ctrl.UserRepo.GetAllUser(c)
}
