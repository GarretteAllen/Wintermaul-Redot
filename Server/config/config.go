package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using defaults")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is required")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	DB = client.Database("WMR")
	if DB == nil {
		log.Fatal("DB is nil after mongo connection")
	}
	log.Println("Connected to MongoDB")
	log.Println("Config initialization complete")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
