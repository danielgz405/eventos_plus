package database

import (
	"context"
	"time"

	"github.com/danielgz405/whale_places/models"
)

func (repo *MongoRepo) AuditOperation(ctx context.Context, user models.Profile, table string, operationType string) error {
	auditLog := models.AuditLog{
		User:  user,
		Type:  operationType,
		Table: table,
		Date:  time.Now(),
	}

	collection := repo.client.Database("whale_places").Collection("audit_logs")
	_, err := collection.InsertOne(ctx, auditLog)
	return err
}
