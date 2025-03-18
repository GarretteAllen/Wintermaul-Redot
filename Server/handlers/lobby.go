package handlers

import (
	"encoding/json"
	"log"

	"wss/services"

	"github.com/gorilla/websocket"
)

func HandleLobby(conn *websocket.Conn, msg []byte) {
	var data struct {
		Action  string `json:"action"`
		LobbyID string `json:"lobby_id"`
	}

	if err := json.Unmarshal(msg, &data); err != nil {
		log.Println("Invalid lobby request")
		return
	}

	switch data.Action {
	case "create":
		lobby := services.CreateLobby()
		conn.WriteJSON(map[string]interface{}{"status": "lobby_created", "lobby": lobby})
	case "join":
		services.JoinLobby(data.LobbyID, conn)
		conn.WriteJSON(map[string]string{"status": "joined_lobby"})
	}
}
