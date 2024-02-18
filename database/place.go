package database

import (
	"context"

	"github.com/danielgz405/whale_places/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertPlace(ctx context.Context, place *models.InsertPlace) (*models.Place, error) {
	collection := repo.client.Database("whale_places").Collection("places")
	result, err := collection.InsertOne(ctx, place)
	if err != nil {
		return nil, err
	}
	createdPlace, err := repo.GetPlaceById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdPlace, nil
}

func (repo *MongoRepo) GetPlaceById(ctx context.Context, id string) (*models.Place, error) {
	collection := repo.client.Database("whale_places").Collection("places")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var place models.Place
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&place)
	if err != nil {
		return nil, err
	}
	return &place, nil
}

func (repo *MongoRepo) ListPlaces(ctx context.Context) ([]models.Place, error) {
	collection := repo.client.Database("whale_places").Collection("places")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var places []models.Place
	if err = cursor.All(ctx, &places); err != nil {
		return nil, err
	}
	return places, nil
}

func (repo *MongoRepo) UpdatePlace(ctx context.Context, data *models.UpdatePlace, id string) (*models.Place, error) {
	collection := repo.client.Database("whale_places").Collection("places")
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
		"coordinates": data.Coordinates,
		"updated_at":  data.UpdatedAt,
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
	updatedPlace, err := repo.GetPlaceById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedPlace, nil
}

func (repo *MongoRepo) DeletePlace(ctx context.Context, id string) error {
	collection := repo.client.Database("whale_places").Collection("places")
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
