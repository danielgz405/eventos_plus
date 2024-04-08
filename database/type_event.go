package database

import (
	"context"

	"github.com/danielgz405/whale_places/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertTypeEvent(ctx context.Context, typeEvent *models.InsertTypeEvent) (*models.TypeEvent, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	result, err := collection.InsertOne(ctx, typeEvent)
	if err != nil {
		return nil, err
	}
	createdTypeEvent, err := repo.GetTypeEventById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdTypeEvent, nil
}

func (repo *MongoRepo) GetTypeEventById(ctx context.Context, id string) (*models.TypeEvent, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var typeEvent models.TypeEvent
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&typeEvent)
	if err != nil {
		return nil, err
	}
	return &typeEvent, nil
}

func (repo *MongoRepo) ListTypeEvents(ctx context.Context) ([]models.TypeEvent, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var typeEvents []models.TypeEvent
	if err = cursor.All(ctx, &typeEvents); err != nil {
		return nil, err
	}
	return typeEvents, nil
}

func (repo *MongoRepo) UpdateTypeEvent(ctx context.Context, data *models.UpdateTypeEvent, id string) (*models.TypeEvent, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"name":        data.Name,
		"description": data.Description,
	}
	for key, value := range iterableData {
		if value != nil && value != "" {
			update["$set"].(bson.M)[key] = value
		}
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return nil, err
	}
	updatedTypeEvent, err := repo.GetTypeEventById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedTypeEvent, nil
}

func (repo *MongoRepo) DeleteTypeEvent(ctx context.Context, id string) error {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
