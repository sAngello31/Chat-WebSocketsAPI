package services

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan string
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*websocket.Conn]bool),
		Broadcast:  make(chan string),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) Run() {
	log.Println("Ejecución del servicio HUB")
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			delete(h.Clients, client)
			client.Close()
		case msg := <-h.Broadcast:
			for conn := range h.Clients {
				err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					conn.Close()
					h.Unregister <- conn
				}
			}
		}
	}
}

func GetUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
}
