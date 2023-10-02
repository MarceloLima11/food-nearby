package geolocation

type Location struct {
	Type        string    `json:"type" bson:"type" validate:"required"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates" validate:"required"`
}
