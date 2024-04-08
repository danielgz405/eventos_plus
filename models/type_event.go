package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TypeEvent struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UserID      string             `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}

type InsertTypeEvent struct {
	UserID      string `json:"user_id" bson:"user_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type UpdateTypeEvent struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
