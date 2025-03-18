package services

import (
	"log"
	"sync"
	"wss/models"
	"wss/utils"

	"github.com/gorilla/websocket"
)

var lobbies = sync.Map{}

func CreateLobby() *models.Lobby {
	// Generate a unique ID for the lobby
	lobbyID := utils.GenerateUniqueID()
	if lobbyID == "" {
		log.Println("Failed to generate unique lobby ID")
		return nil
	}

	lobby := &models.Lobby{
		ID:      lobbyID,
		Players: []string{},
		MaxSize: 9,
	}

	lobbies.Store(lobbyID, lobby)
	log.Println("Lobby created with ID:", lobbyID)
	return lobby
}

func JoinLobby(lobbyID string, conn *websocket.Conn, username string) {
	lobbyInterface, exists := lobbies.Load(lobbyID)
	if !exists {
		log.Println("Lobby not found:", lobbyID)
		conn.WriteJSON(map[string]string{"status": "error", "message": "Lobby not found"})
		return
	}

	lobby := lobbyInterface.(*models.Lobby)

	if len(lobby.Players) >= lobby.MaxSize {
		log.Println("Lobby full:", lobbyID)
		conn.WriteJSON(map[string]string{"status": "error", "message": "Lobby is full"})
		return
	}

	// TODO: Replace "newPlayer" with the actual player's username from the game message
	lobby.Players = append(lobby.Players, username)

	lobbies.Store(lobbyID, lobby)
	log.Println("Player", username, "joined lobby:", lobbyID)
}

func DeleteLobby(lobbyID string) {
	lobbies.Delete(lobbyID)
	log.Println("Lobby deleted:", lobbyID)
}
