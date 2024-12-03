package controllers

import (
	"chat_websocket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitChat(c *gin.Context) {
	/*
		msgString := ("Conexion WebSocket. ID: " + token)
		ws, err := services.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Error al conectar el websocket: ", err)
			c.JSON(http.StatusBadRequest, "Error al conectar el websocket")
			return
		}
		defer ws.Close()
		for {

			ws.WriteMessage(websocket.TextMessage, []byte(msgString))
			time.Sleep(time.Second)
		}
	*/
	s := services.GenerateUUID(32, 21342)
	c.JSON(http.StatusOK, gin.H{"message": s})
}
