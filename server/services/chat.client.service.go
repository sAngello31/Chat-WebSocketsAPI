package services

import (
	"chat_websocket/models"
	"chat_websocket/utils"
	"context"
	"log"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
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
		msg.UUID = c.UUID
		c.Hub.Broadcast <- msg
	}
}

func (c *ChatClient) CloseClient() {
	c.Hub.Unregister <- c
}

func (c *ChatClient) SendMsg(msg *models.Message) error {
	err := c.Conn.WriteJSON(msg)
	if err != nil {
		return err
	}
	return err
}

func saveMsg(msg *models.Message) (*mongo.InsertOneResult, error) {
	collection := utils.GetCollection("chats")
	return collection.InsertOne(context.TODO(), msg)
}
