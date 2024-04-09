package database

import (
	"context"
	"fmt"

	"github.com/danielgz405/whale_places/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *MongoRepo) InsertEvent(ctx context.Context, event *models.InsertEvent) (*models.Event, error) {
	collection := repo.client.Database("whale_places").Collection("event")
	result, err := collection.InsertOne(ctx, event)
	if err != nil {
		return nil, err
	}
	createdEvent, err := repo.GetEventById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdEvent, nil
}

func (repo *MongoRepo) GetEventById(ctx context.Context, id string) (*models.Event, error) {
	collection := repo.client.Database("whale_places").Collection("event")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var event models.Event
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (repo *MongoRepo) ListEvents(ctx context.Context) ([]models.Event, error) {
	collection := repo.client.Database("whale_places").Collection("event")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	events := []models.Event{}

	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}

func (repo *MongoRepo) UpdateEvent(ctx context.Context, data *models.UpdateEvent, id string) (*models.Event, error) {
	collection := repo.client.Database("whale_places").Collection("event")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"type_event": data.TypeEvent,
		"name":       data.Name,
		"is_free":    data.IsFree,
		"capacity":   data.Capacity,
		"type_tiket": data.TypeTiket,
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
	updatedEvent, err := repo.GetEventById(ctx, id)
	if err != nil {
		return nil, err
	}
	return updatedEvent, nil
}

func (repo *MongoRepo) DeleteEvent(ctx context.Context, id string) error {
	collection := repo.client.Database("whale_places").Collection("event")

	event, err := repo.GetEventById(ctx, id)
	if err != nil {
		return err
	}

	reserve, err := repo.ListReservesByEvent(ctx, event.ID.Hex())
	if err != nil {
		return err
	}

	if len(reserve) > 0 {
		return fmt.Errorf("This event has users reserve")
	}

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

func (repo *MongoRepo) ListEventsByPage(ctx context.Context, limit int, page int) ([]models.Event, int, error) {
	collection := repo.client.Database("whale_places").Collection("event")
	offset := (page - 1) * limit

	quantity, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}

	events := []models.Event{}
	if err = cursor.All(ctx, &events); err != nil {
		return nil, 0, err
	}

	return events, int(quantity), nil
}

func (repo *MongoRepo) ListEventsByName(ctx context.Context, limit int, page int, name string) ([]models.Event, int, error) {
	collection := repo.client.Database("whale_places").Collection("event")
	offset := (page - 1) * limit

	filter := bson.M{"name": primitive.Regex{Pattern: name, Options: "i"}}

	quantity, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Find(ctx, filter, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}
	events := []models.Event{}
	if err = cursor.All(ctx, &events); err != nil {
		return nil, 0, err
	}

	return events, int(quantity), nil
}
