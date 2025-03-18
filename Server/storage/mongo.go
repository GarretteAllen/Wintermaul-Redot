package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var Database *mongo.Database

func GetCollection(name string) *mongo.Collection {
	return Database.Collection(name)
}
