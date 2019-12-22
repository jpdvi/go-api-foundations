package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jpdvi/chat-service/models"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love beans!")
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWsConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error: %v", err)
	}
	defer ws.Close()

	clients[ws] = true
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
		}

		broadcast <- msg
	}

}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {

			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}

			log.Println(msg)
		}
	}
}
