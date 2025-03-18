package services

import (
	"log"
	"wss/models"

	"github.com/gorilla/websocket"
)

var lobbies = make(map[string]*models.Lobby)

// TODO: get lobby name from godot message to append lobby players.
func CreateLobby() *models.Lobby {
	lobby := &models.Lobby{
		ID:      "randomID",
		Players: []string{},
		MaxSize: 9,
	}
	lobbies[lobby.ID] = lobby
	log.Println("Lobby created:", lobby.ID)
	return lobby
}

func JoinLobby(lobbyID string, conn *websocket.Conn) {
	lobby, exists := lobbies[lobbyID]
	if !exists {
		log.Println("Lobby not found:", lobbyID)
		return
	}
	if len(lobby.Players) >= lobby.MaxSize {
		log.Println("Lobby full:", lobbyID)
		return
	}
	// TODO: get player username from godot message to append lobby players.
	lobby.Players = append(lobby.Players, "newPlayer")
	log.Println("Player joined lobby:", lobbyID)
}
