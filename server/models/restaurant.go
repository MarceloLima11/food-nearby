package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Restaurant struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name, omitempty" bson:"name" validate:"required"`
	CuisineType string             `json:"cuisine-type, omitempty" bson:"address" validate:"required"`
	Location    []float64          `json:"location" bson:"location" validate:"required"`
	About       string             `json:"about, omitempty" bson:"about" validate:"required"`
}
