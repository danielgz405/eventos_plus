package repository

import (
	"context"

	"github.com/danielgz405/whale_places/models"
)

func InsertTypeEvent(ctx context.Context, typeEvent *models.InsertTypeEvent) (*models.TypeEvent, error) {
	return implementation.InsertTypeEvent(ctx, typeEvent)
}

func GetTypeEventById(ctx context.Context, id string) (*models.TypeEvent, error) {
	return implementation.GetTypeEventById(ctx, id)
}

func ListTypeEvents(ctx context.Context) ([]models.TypeEvent, error) {
	return implementation.ListTypeEvents(ctx)
}

func DeleteTypeEvent(ctx context.Context, id string) error {
	return implementation.DeleteTypeEvent(ctx, id)
}

func UpdateTypeEvent(ctx context.Context, data *models.UpdateTypeEvent, id string) (*models.TypeEvent, error) {
	return implementation.UpdateTypeEvent(ctx, data, id)
}
