package handlers

import (
	"encoding/json"
	"log"
	"wss/constants"
	"wss/services"
	"wss/utils"

	"github.com/gorilla/websocket"
)

func HandleLobby(conn *websocket.Conn, msg []byte) {
	var data struct {
		Action  string `json:"action"`
		LobbyID string `json:"lobby_id"`
		Token   string `json:"token"`
	}

	if err := json.Unmarshal(msg, &data); err != nil {
		log.Println("Invalid lobby request")
		return
	}

	switch data.Action {
	case constants.ActionCreateLobby:
		lobby := services.CreateLobby()
		conn.WriteJSON(map[string]interface{}{"status": constants.StatusLobbyCreated, "lobby": lobby})
	case constants.ActionJoinLobby:
		username, err := utils.ValidateToken(data.Token)
		if err != nil {
			log.Println("Token validation falied", err)
			conn.WriteJSON(map[string]string{"status": "error", "message": "Invalid token"})
			return
		}
		services.JoinLobby(data.LobbyID, conn, username)
		conn.WriteJSON(map[string]string{"status": constants.StatusJoinedLobby})
	default:
		log.Println("Unknown action:", data.Action)
	}
}
