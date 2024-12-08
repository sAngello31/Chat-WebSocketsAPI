package controllers

import (
	"chat_websocket/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	newClient := services.NewClient(c.Param("uuid"), conn)
	log.Println(newClient.Hub.Clients)
	go newClient.ReadMsg()

}
