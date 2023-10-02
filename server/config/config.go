package config

import (
	"github.com/MarceloLima11/food-nearby/server/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db *mongo.Collection
)

func Init() error {
	var err error

	db, err = initializeMongo()
	utils.IfErrThrowFatalf(err, "Error initializing mongodb: ")

	return nil
}

func GetMongoDb() *mongo.Collection {
	return db
}
