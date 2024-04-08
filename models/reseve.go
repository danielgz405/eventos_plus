package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reserve struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	UserID        string             `json:"user_id" bson:"user_id"`
	EventId       string             `json:"event_id" bson:"event_id"`
	Acceted       bool               `json:"acceted" bson:"acceted"`
	DateToAcceted time.Time          `json:"date_to_acceted" bson:"date_to_acceted"`
}

type InsertReserve struct {
	UserID        string    `json:"user_id" bson:"user_id"`
	EventId       string    `json:"event_id" bson:"event_id"`
	Acceted       bool      `json:"acceted" bson:"acceted"`
	DateToAcceted time.Time `json:"date_to_acceted" bson:"date_to_acceted"`
}

type UpdateReserve struct {
	Acceted       bool      `json:"acceted" bson:"acceted"`
	DateToAcceted time.Time `json:"date_to_acceted" bson:"date_to_acceted"`
}
