package services

import "github.com/gorilla/websocket"

type ChatClient struct {
	Hub  *Hub
	Conn *websocket.Conn
	UUID string
	Send chan []byte
}

func NewClient(uuid string, conn *websocket.Conn) *ChatClient {
	return &ChatClient{
		Hub:  GetHub(),
		Conn: conn,
		UUID: uuid,
		Send: make(chan []byte),
	}
}

// Revisar cual es el mejor
func (c *ChatClient) CloseClient() {
	c.Hub.Unregister <- c
	c.Conn.Close()
}
