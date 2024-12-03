package controllers

import (
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
}
