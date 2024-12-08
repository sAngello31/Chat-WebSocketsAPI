package services

import (
	"bytes"
	"log"

	"github.com/gorilla/websocket"
)

type ChatClient struct {
	Hub  *Hub
	Conn *websocket.Conn
	UUID string
	Send chan []byte
}

func NewClient(uuid string, conn *websocket.Conn) *ChatClient {
	client := ChatClient{
		Hub:  GetHub(),
		Conn: conn,
		UUID: uuid,
		Send: make(chan []byte),
	}
	client.Hub.Register <- &client
	return &client

}

func (c *ChatClient) ReadMsg() {
	defer c.CloseClient()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		msg = bytes.TrimSpace(bytes.Replace(msg, []byte("\n"), []byte(" "), -1))
		c.Hub.Broadcast <- msg
	}
}

func (c *ChatClient) CloseClient() {
	c.Hub.Unregister <- c
}
