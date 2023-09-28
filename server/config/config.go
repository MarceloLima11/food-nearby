package config

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db *mongo.Collection
)

func Init() error {
	var err error

	db, err = initializeMongo()

	if err != nil {
		return fmt.Errorf("Error initializing mongodb: %v", err)
	}

	return nil
}

func GetMongoDb() *mongo.Collection {
	return db
}
