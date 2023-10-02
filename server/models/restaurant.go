package model

import (
	"github.com/MarceloLima11/food-nearby/server/geolocation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	Id          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name, omitempty" bson:"name" validate:"required"`
	CuisineType string               `json:"cuisineType, omitempty" bson:"cuisineType" validate:"required"`
	About       string               `json:"about, omitempty" bson:"about" validate:"required"`
	Location    geolocation.Location `json:"location" bson:"location" validate:"required"`
}
