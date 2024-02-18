package structures

import (
	"github.com/danielgz405/whale_places/models"
)

type InsertPlaceRequest struct {
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Coordinates models.Coordinates `json:"coordinates" bson:"coordinates"`
}

type UpdatePlaceRequest struct {
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Coordinates models.Coordinates `json:"coordinates" bson:"coordinates"`
}
