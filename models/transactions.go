package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	RecipientId string             `json:"recipient_id" bson:"recipient_id"`
	EmmiterId   string             `json:"emiter_id" bson:"emiter_id"`
	EventId     string             `json:"event_id" bson:"event_id"`
	ReserveId   string             `json:"reserve_id" bson:"reserve_id"`
	CreateAt    time.Time          `json:"create_at" bson:"create_at"`
	Total       float64            `json:"total" bson:"total"`
	Ref         string             `json:"ref" bson:"ref"`
	Subtotal    float64            `json:"subtotal" bson:"subtotal"`
}

type InsertTransaction struct {
	RecipientId string    `json:"recipient_id" bson:"recipient_id"`
	EmmiterId   string    `json:"emiter_id" bson:"emiter_id"`
	EventId     string    `json:"event_id" bson:"event_id"`
	ReserveId   string    `json:"reserve_id" bson:"reserve_id"`
	CreateAt    time.Time `json:"create_at" bson:"create_at"`
	Total       float64   `json:"total" bson:"total"`
	Ref         string    `json:"ref" bson:"ref"`
	Subtotal    float64   `json:"subtotal" bson:"subtotal"`
}
