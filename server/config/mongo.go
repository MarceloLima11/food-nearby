package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"
const dbName = "food-nearby"

func initializeMongo() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection Done")

	collection := client.Database(dbName).Collection("restaurant")
	fmt.Println("Collection is ready")

	return collection, nil
}
