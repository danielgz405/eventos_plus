package database

import (
	"context"

	"github.com/danielgz405/whale_places/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertReserve(ctx context.Context, typeEvent *models.InsertReserve) (*models.Reserve, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	result, err := collection.InsertOne(ctx, typeEvent)
	if err != nil {
		return nil, err
	}
	createdReserve, err := repo.GetReserveById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdReserve, nil
}

func (repo *MongoRepo) GetReserveById(ctx context.Context, id string) (*models.Reserve, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var typeEvent models.Reserve
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&typeEvent)
	if err != nil {
		return nil, err
	}
	return &typeEvent, nil
}

func (repo *MongoRepo) ListReservesByUser(ctx context.Context, userId string) ([]models.Reserve, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	cursor, err := collection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	var typeEvents []models.Reserve
	if err = cursor.All(ctx, &typeEvents); err != nil {
		return nil, err
	}
	return typeEvents, nil
}

func (repo *MongoRepo) ListReservesByEvent(ctx context.Context, eventId string) ([]models.Reserve, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	cursor, err := collection.Find(ctx, bson.M{"event_id": eventId})
	if err != nil {
		return nil, err
	}
	var typeEvents []models.Reserve
	if err = cursor.All(ctx, &typeEvents); err != nil {
		return nil, err
	}
	return typeEvents, nil
}

func (repo *MongoRepo) UpdateReserve(ctx context.Context, data *models.UpdateReserve, id string) (*models.Reserve, error) {
	collection := repo.client.Database("whale_places").Collection("typeEvents")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"acceted":         data.Acceted,
		"date_to_acceted": data.DateToAcceted,
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
	updatedReserve, err := repo.GetReserveById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedReserve, nil
}
