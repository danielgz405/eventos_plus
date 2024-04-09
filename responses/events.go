package responses

import "github.com/danielgz405/whale_places/models"

type EventsResponse struct {
	Quantity int            `json:"quantity"`
	Event    []models.Event `json:"events"`
}
