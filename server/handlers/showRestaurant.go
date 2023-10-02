package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	model "github.com/MarceloLima11/food-nearby/server/models"
	"github.com/MarceloLima11/food-nearby/server/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ShowRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantId := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(restaurantId)
	if utils.IfErrThrowWriteError(err, w, utils.InvalidID, http.StatusBadRequest) {
		return
	}

	filter := bson.M{"_id": objectID}

	var restaurant model.Restaurant
	err = collection.FindOne(context.Background(), filter).Decode(&restaurant)
	if utils.IfErrThrowWriteError(err, w, utils.DataNotFound, http.StatusNotFound) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurant)
}
