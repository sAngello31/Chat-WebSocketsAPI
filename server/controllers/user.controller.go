package controllers

import (
	"chat_websocket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	services.GetAllUsers(c)
}

func GetUserData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"username":       c.Keys["username"],
		"contact_number": c.Keys["contact_number"],
	})
}
