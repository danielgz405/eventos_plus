package repository

import (
	"context"

	"github.com/danielgz405/whale_places/models"
)

func InsertReserve(ctx context.Context, reserve *models.InsertReserve) (*models.Reserve, error) {
	return implementation.InsertReserve(ctx, reserve)
}

func GetReserveById(ctx context.Context, id string) (*models.Reserve, error) {
	return implementation.GetReserveById(ctx, id)
}

func ListReservesByUser(ctx context.Context, userId string) ([]models.Reserve, error) {
	return implementation.ListReservesByUser(ctx, userId)
}

func UpdateReserve(ctx context.Context, data *models.UpdateReserve, id string) (*models.Reserve, error) {
	return implementation.UpdateReserve(ctx, data, id)
}
