package main

import (
	"fmt"
	"log"
	"net/http"

	"wss/config"
	"wss/handlers"

	"github.com/gorilla/mux"
)

func main() {
	config.InitConfig()

	router := mux.NewRouter()
	router.HandleFunc("/ws", handlers.HandleConnections)
	port := ":8080"
	fmt.Println("Server running on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
