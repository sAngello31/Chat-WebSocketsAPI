package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func GetUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
}

func HandleConn(c *gin.Context) {
	upgrader := GetUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error: Websocket")
		return
	}
	newClient := NewClient(c.Param("uuid"), conn)
	log.Println(newClient.Hub.Clients)
	go newClient.ReadMsg()
}
