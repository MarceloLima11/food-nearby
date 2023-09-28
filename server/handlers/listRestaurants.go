package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MarceloLima11/food-nearby/server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	longitude, err := strconv.ParseFloat(r.URL.Query().Get("long"), 64)
	utils.CheckNilError(err)

	latitude, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	utils.CheckNilError(err)

	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longitude, latitude},
				},
				"$maxDistance": 5000,
			},
		},
	}

	cur, err := collection.Find(context.Background(), filter)
	utils.CheckNilError(err)

	var restaurants []primitive.M
	for cur.Next(context.Background()) {
		var restaurant bson.M
		err := cur.Decode(&restaurant)
		utils.CheckNilError(err)
		restaurants = append(restaurants, restaurant)
	}

	defer cur.Close(context.Background())
	json.NewEncoder(w).Encode(restaurants)
}
