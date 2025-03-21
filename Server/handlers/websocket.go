package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = sync.Map{}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming WebSocket connection attempt")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket Upgrade failed:", err)
		return
	}
	defer func() {
		// disconnect cleanup
		clients.Delete(conn)
		conn.Close()
		log.Println("Connection closed")
	}()

	clients.Store(conn, true)
	log.Println("New WebSocket connection established")

	go startHeartbeat(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Websocket read error:", err)
			break
		}
		log.Printf("Received: %s\n", msg)

		var message struct {
			Action string `json:"action"`
		}
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}
		switch message.Action {
		case "register", "login":
			HandleAuth(conn, msg)
		case "create_lobby", "join_lobby":
			HandleLobby(conn, msg)
		default:
			log.Print("Unknown action:", message.Action)
		}
	}
}

// send message to all active clients
func BroadcastMessage(message []byte) {
	clients.Range(func(key, value interface{}) bool {
		conn, ok := key.(*websocket.Conn)
		if !ok {
			return false
		}
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error broadcasting message", err)
			clients.Delete(conn)
			conn.Close()
		}
		return true
	})
}

// check connection health with ticker
func startHeartbeat(conn *websocket.Conn) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Println("Heartbeat failed for connection:", err)
				clients.Delete(conn)
				conn.Close()
				return
			}
		}
	}
}
