package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	model "github.com/MarceloLima11/food-nearby/server/models"
	"github.com/MarceloLima11/food-nearby/server/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantId := vars["id"]

	var updatedRestaurant model.Restaurant
	err := json.NewDecoder(r.Body).Decode(&updatedRestaurant)
	utils.CheckNilError(err)

	validate := validator.New()
	err = validate.Struct(updatedRestaurant)
	utils.CheckNilError(err)

	filter := bson.M{"_id": restaurantId}
	update := bson.M{"$set": updatedRestaurant}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	utils.CheckNilError(err)

	if result.ModifiedCount == 0 {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
