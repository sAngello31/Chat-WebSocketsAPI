package controllers

import (
	"chat_websocket/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	services.LoginUser(c)
}

func RegisterNewUser(c *gin.Context) {
	services.RegisterUser(c)
}
