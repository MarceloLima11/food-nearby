package handlers

import (
	"github.com/MarceloLima11/food-nearby/server/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
)

func InitializeHandler() {
	collection = config.GetMongoDb()
}
