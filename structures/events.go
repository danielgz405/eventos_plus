package structures

import "github.com/danielgz405/whale_places/models"

type InsertEventsRequest struct {
	PlaceID     string               `json:"place_id" bson:"place_id"`
	TypeEvent   string               `json:"type_event" bson:"type_event"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Date        string               `json:"date" bson:"date"`
	IsFree      bool                 `json:"is_free" bson:"is_free"`
	Capacity    int64                `json:"capacity" bson:"capacity"`
	TypeTiket   []models.TypesTikets `json:"type_tiket" bson:"type_tiket"`
}

type UpdateEventsRequest struct {
	TypeEvent string               `json:"type_event" bson:"type_event"`
	Name      string               `json:"name" bson:"name"`
	IsFree    bool                 `json:"is_free" bson:"is_free"`
	Capacity  int64                `json:"capacity" bson:"capacity"`
	TypeTiket []models.TypesTikets `json:"type_tiket" bson:"type_tiket"`
}
