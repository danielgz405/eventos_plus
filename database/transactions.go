package database

import (
	"context"

	"github.com/danielgz405/whale_places/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertTransaction(ctx context.Context, transaction *models.InsertTransaction) (*models.Transaction, error) {
	collection := repo.client.Database("whale_places").Collection("transactions")
	result, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		return nil, err
	}
	createdTransaction, err := repo.GetTransactionById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdTransaction, nil
}

func (repo *MongoRepo) GetTransactionById(ctx context.Context, id string) (*models.Transaction, error) {
	collection := repo.client.Database("whale_places").Collection("transactions")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var transaction models.Transaction
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (repo *MongoRepo) ListTransactionsByUser(ctx context.Context, userId string) ([]models.Transaction, error) {
	collection := repo.client.Database("whale_places").Collection("transactions")
	cursor, err := collection.Find(ctx, bson.M{
		"$or": []bson.M{
			{"recipient_id": userId},
			{"emiter_id": userId},
		},
	})
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}
