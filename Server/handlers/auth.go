package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"wss/models"
	"wss/storage"
	"wss/utils"

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
		// hash password
		hashedPassword, err := utils.HashPassword(data.Password)
		if err != nil {
			conn.WriteJSON(map[string]string{"status": "error", "message": "Internal server error"})
			return
		}

		// create new user
		newUser := models.User{
			Username: data.Username,
			Password: hashedPassword,
			Created:  time.Now(),
		}

		_, err = usersCollection.InsertOne(context.TODO(), newUser)
		if err != nil {
			log.Println("Registration failed:", err)
			conn.WriteJSON(map[string]string{"status": "error", "message": "Registration failed"})
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
			conn.WriteJSON(map[string]string{"status": "error", "message": "Internal server error"})
			return
		}

		// verify password
		if !utils.VerifyPassword(data.Password, user.Password) {
			conn.WriteJSON(map[string]string{"status": "error", "message": "Invalid password"})
			return
		}

		// generate token
		token, err := utils.GenerateToken(user.Username)
		if err != nil {
			log.Println("Token generation error:", err)
			conn.WriteJSON(map[string]string{"status": "error", "message": "Internal server error"})
			return
		}

		conn.WriteJSON(map[string]string{"status": "logged_in", "token": token})
	}
}
