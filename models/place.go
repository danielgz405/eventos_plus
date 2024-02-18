package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Place struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Coordinates Coordinates        `json:"coordinates" bson:"coordinates"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type InsertPlace struct {
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates"`
	CreatedAt   time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at"`
}

type UpdatePlace struct {
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates"`
	UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}
