package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket Upgrade failed:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	log.Println("New websocket connection")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Websocket read error:", err)
			delete(clients, conn)
			break
		}
		log.Printf("Received: %s\n", msg)
	}
}
