package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"wss/models"
	"wss/storage"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleAuth(conn *websocket.Conn, msg []byte) {
	var data struct {
		Action   string `json:"action"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.Unmarshal(msg, &data); err != nil {
		log.Println("Invalid auth request")
		return
	}

	usersCollection := storage.GetCollection("users")

	if data.Action == "register" {
		newUser := models.User{
			Username: data.Username,
			Password: data.Password,
			Created:  time.Now(),
		}

		_, err := usersCollection.InsertOne(context.TODO(), newUser)
		if err != nil {
			log.Println("Registration failed:", err)
			return
		}
		conn.WriteJSON(map[string]string{"status": "registered"})
	} else if data.Action == "login" {
		var user models.User
		err := usersCollection.FindOne(context.TODO(), bson.M{"username": data.Username}).Decode(&user)
		if err == mongo.ErrNoDocuments {
			conn.WriteJSON(map[string]string{"status": "error", "message": "User not found"})
			return
		} else if err != nil {
			log.Println("Login error:", err)
			return
		}
		conn.WriteJSON(map[string]string{"status": "logged_in"})
	}
}
