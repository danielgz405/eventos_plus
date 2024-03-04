package repository

import (
	"context"

	"github.com/danielgz405/whale_places/models"
)

func InsertPlace(ctx context.Context, place *models.InsertPlace) (*models.Place, error) {
	return implementation.InsertPlace(ctx, place)
}

func GetPlaceById(ctx context.Context, id string) (*models.Place, error) {
	return implementation.GetPlaceById(ctx, id)
}

func ListPlaces(ctx context.Context) ([]models.Place, error) {
	return implementation.ListPlaces(ctx)
}

func DeletePlace(ctx context.Context, id string) error {
	return implementation.DeletePlace(ctx, id)
}

func UpdatePlace(ctx context.Context, data *models.UpdatePlace, id string) (*models.Place, error) {
	return implementation.UpdatePlace(ctx, data, id)
}
