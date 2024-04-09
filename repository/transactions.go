package repository

import (
	"context"

	"github.com/danielgz405/whale_places/models"
)

func InsertTransaction(ctx context.Context, reserve *models.InsertTransaction) (*models.Transaction, error) {
	return implementation.InsertTransaction(ctx, reserve)
}

func GetTransactionById(ctx context.Context, id string) (*models.Transaction, error) {
	return implementation.GetTransactionById(ctx, id)
}

func ListTransactionsByUser(ctx context.Context, userId string) ([]models.Transaction, error) {
	return implementation.ListTransactionsByUser(ctx, userId)
}
