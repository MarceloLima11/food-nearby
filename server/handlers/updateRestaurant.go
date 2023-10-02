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
	if utils.IfErrThrowWriteError(err, w, utils.DecodeJSON, http.StatusBadRequest) {
		return
	}

	validate := validator.New()
	err = validate.Struct(updatedRestaurant)
	if utils.IfErrThrowWriteError(err, w, utils.Validation, http.StatusBadRequest) {
		return
	}

	filter := bson.M{"_id": restaurantId}
	update := bson.M{"$set": updatedRestaurant}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if utils.IfErrThrowWriteError(err, w, utils.Internal, http.StatusInternalServerError) {
		return
	}

	if result.ModifiedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(utils.DataNotFound))
		return
	}

	w.WriteHeader(http.StatusOK)
}
