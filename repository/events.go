package repository

import (
	"context"

	"github.com/danielgz405/whale_places/models"
)

func InsertEvent(ctx context.Context, event *models.InsertEvent) (*models.Event, error) {
	return implementation.InsertEvent(ctx, event)
}

func GetEventById(ctx context.Context, id string) (*models.Event, error) {
	return implementation.GetEventById(ctx, id)
}

func ListEvents(ctx context.Context) ([]models.Event, error) {
	return implementation.ListEvents(ctx)
}

func UpdateEvent(ctx context.Context, data *models.UpdateEvent, id string) (*models.Event, error) {
	return implementation.UpdateEvent(ctx, data, id)
}

func DeleteEvent(ctx context.Context, id string) error {
	return implementation.DeleteEvent(ctx, id)
}

func ListEventsByPage(ctx context.Context, limit int, page int) ([]models.Event, int, error) {
	return implementation.ListEventsByPage(ctx, limit, page)
}

func ListEventsByName(ctx context.Context, limit int, page int, name string) ([]models.Event, int, error) {
	return implementation.ListEventsByName(ctx, limit, page, name)
}
