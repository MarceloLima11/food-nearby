package handlers

import (
	"context"
	"net/http"

	"github.com/MarceloLima11/food-nearby/server/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	restaurantId := params["id"]
	id, err := primitive.ObjectIDFromHex(restaurantId)
	if utils.IfErrThrowWriteError(err, w, utils.InvalidID, http.StatusBadRequest) {
		return
	}

	filter := bson.M{"_id": id}
	delCount, err := collection.DeleteOne(context.Background(), filter)
	if utils.IfErrThrowWriteError(err, w, utils.DeleteData,
		http.StatusInternalServerError) {
		return
	}

	if delCount.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(utils.DeleteDataNotFound))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(utils.DeleteSuccess))
}
