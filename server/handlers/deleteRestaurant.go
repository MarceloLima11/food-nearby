package handlers

import (
	"context"
	"fmt"
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
	utils.CheckNilError(err)

	filter := bson.M{"_id": id}
	delCount, err := collection.DeleteOne(context.Background(), filter)
	utils.CheckNilError(err)

	w.WriteHeader(http.StatusNoContent)

	response := fmt.Sprintf("Deleted register: %d", delCount.DeletedCount)
	w.Write([]byte(response))
}
