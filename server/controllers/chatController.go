package controllers

import (
	"chat_websocket/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func GetUUIDUsers(c *gin.Context) {

	userA := c.Param("userA")
	userB := c.Param("userB")

	if userA == "" || userB == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	s := services.GenerateUUID(userA, userB)
	c.JSON(http.StatusOK, gin.H{"message": s})
}

func HandleConn(c *gin.Context) {
	upgrader := services.GetUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error: Websocket")
		return
	}
	defer conn.Close()
	s := fmt.Sprint(c.Keys["username"], " se ha conectado al chat")
	conn.WriteMessage(websocket.TextMessage, []byte(s))

}
