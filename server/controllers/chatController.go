package controllers

import (
	"chat_websocket/models"
	"chat_websocket/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatController struct {
	MsgRepo *models.MsgRepository
}

var ThisChatController *ChatController

func NewChatController(msgRepo *models.MsgRepository) {
	ThisChatController = &ChatController{MsgRepo: msgRepo}
}

func (ctrl *ChatController) InitChat(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	token, err := services.GetIDFromJWT(tokenString)
	if err != nil {
		log.Println("Acceso no Autorizado", err)
		c.JSON(http.StatusBadRequest, "Acceso no Autorizado")
		return
	}
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
}
