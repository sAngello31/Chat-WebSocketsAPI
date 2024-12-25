package services

import (
	"chat_websocket/models"
	"log"
)

var hub *Hub

type Hub struct {
	Clients    map[*ChatClient]bool
	Broadcast  chan models.Message
	Register   chan *ChatClient
	Unregister chan *ChatClient
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*ChatClient]bool),
		Broadcast:  make(chan models.Message),
		Register:   make(chan *ChatClient),
		Unregister: make(chan *ChatClient),
	}
}

func (h *Hub) Run() {
	log.Println("Ejecución del servicio Hub")
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			delete(h.Clients, client)
			client.Conn.Close()
		case msg := <-h.Broadcast:
			h.sendMsg(&msg)
		}
	}
}

func (h *Hub) sendMsg(msg *models.Message) {
	for client := range h.Clients {
		if client.UUID != msg.UUID {
			continue
		}
		err := client.SendMsg(msg)
		if err != nil {
			log.Println(err)
			h.Unregister <- client
		}
	}
	if _, err := saveMsg(msg); err != nil {
		log.Println(err)
	}
}

func InitHub() {
	hub = NewHub()
	hub.Run()
}

func GetHub() *Hub {
	return hub
}
