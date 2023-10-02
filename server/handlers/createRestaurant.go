package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/MarceloLima11/food-nearby/server/models"
	"github.com/MarceloLima11/food-nearby/server/utils"

	"github.com/go-playground/validator/v10"
)

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var restaurant model.Restaurant
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if utils.IfErrThrowWriteError(err, w, utils.DecodeJSON, http.StatusBadRequest) {
		return
	}

	validate := validator.New()

	err = validate.Struct(restaurant)
	if utils.IfErrThrowWriteError(err, w, utils.Validation, http.StatusBadRequest) {
		return
	}

	inserted, err := collection.InsertOne(context.Background(), restaurant)
	if utils.IfErrThrowWriteError(err, w, utils.InsertRestaurant, http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("Inserted one restaurant with ID:%v", inserted.InsertedID)
}
