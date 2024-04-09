package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	PlaceID     string             `json:"place_id" bson:"place_id"`
	TypeEvent   string             `json:"type_event" bson:"type_event"`
	UserID      string             `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Date        time.Time          `json:"date" bson:"date"`
	IsFree      bool               `json:"is_free" bson:"is_free"`
	Capacity    int64              `json:"capacity" bson:"capacity"`
	TypeTiket   []TypesTikets      `json:"type_tiket" bson:"type_tiket"`
}

type InsertEvent struct {
	PlaceID     string        `json:"place_id" bson:"place_id"`
	TypeEvent   string        `json:"type_event" bson:"type_event"`
	UserID      string        `json:"user_id" bson:"user_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Date        time.Time     `json:"date" bson:"date"`
	IsFree      bool          `json:"is_free" bson:"is_free"`
	Capacity    int64         `json:"capacity" bson:"capacity"`
	TypeTiket   []TypesTikets `json:"type_tiket" bson:"type_tiket"`
}

type UpdateEvent struct {
	TypeEvent string        `json:"type_event" bson:"type_event"`
	Name      string        `json:"name" bson:"name"`
	IsFree    bool          `json:"is_free" bson:"is_free"`
	Capacity  int64         `json:"capacity" bson:"capacity"`
	TypeTiket []TypesTikets `json:"type_tiket" bson:"type_tiket"`
}

type TypesTikets struct {
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}
