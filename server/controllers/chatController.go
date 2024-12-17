package controllers

import (
	"chat_websocket/services"
	"log"

	"github.com/gin-gonic/gin"
)

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
