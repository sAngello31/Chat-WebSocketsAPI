package controllers

import (
	"chat_websocket/services"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	services.GetAllUsers(c)
}
