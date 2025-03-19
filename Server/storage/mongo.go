package storage

import (
	"wss/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var Database *mongo.Database

func GetCollection(name string) *mongo.Collection {
	return config.DB.Collection(name)
}
