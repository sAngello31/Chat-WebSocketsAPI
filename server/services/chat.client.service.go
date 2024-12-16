package services

import (
	"chat_websocket/models"
	"log"

	"github.com/gorilla/websocket"
)

type ChatClient struct {
	Hub  *Hub
	Conn *websocket.Conn
	UUID string
	Send chan models.Message
}

func NewClient(uuid string, conn *websocket.Conn) *ChatClient {
	client := ChatClient{
		Hub:  GetHub(),
		Conn: conn,
		UUID: uuid,
		Send: make(chan models.Message),
	}
	client.Hub.Register <- &client
	return &client

}

func (c *ChatClient) ReadMsg() {
	defer c.CloseClient()
	for {
		var msg models.Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}
		c.Hub.Broadcast <- msg
	}
}

func (c *ChatClient) CloseClient() {
	c.Hub.Unregister <- c
}
