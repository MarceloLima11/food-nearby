package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/MarceloLima11/food-nearby/server/models"
	"github.com/MarceloLima11/food-nearby/server/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func ListRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	longitude, err := strconv.ParseFloat(r.URL.Query().Get("long"), 64)
	if utils.IfErrThrowWriteError(err, w, utils.InvalidCoordinatesParam, http.StatusBadRequest) {
		return
	}

	latitude, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if utils.IfErrThrowWriteError(err, w, utils.InvalidCoordinatesParam, http.StatusBadRequest) {
		return
	}

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
	if utils.IfErrThrowWriteError(err, w, utils.Internal,
		http.StatusInternalServerError) {
		return
	}
	defer cur.Close(context.Background())

	var restaurants []model.Restaurant
	for cur.Next(context.Background()) {
		var restaurant model.Restaurant
		err := cur.Decode(&restaurant)
		utils.IfErrThrowWriteError(err, w, utils.Internal, http.StatusInternalServerError)
		restaurants = append(restaurants, restaurant)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}
