package repository

import (
	"context"

	"github.com/danielgz405/whale_places/models"
)

func AuditOperation(ctx context.Context, user models.Profile, table string, operationType string) error {
	return implementation.AuditOperation(ctx, user, table, operationType)
}
