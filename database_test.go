package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/danielgz405/whale_places/database"
	"github.com/danielgz405/whale_places/models"
	"github.com/joho/godotenv"
)

func TestInsertPlace(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	repo, err := database.NewMongoRepo(os.Getenv("DB_URI_TEST"))
	if err != nil {
		repo.Close()
		log.Fatal(err)
		return
	}

	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		t.Errorf("error loading location: %v", err)
		repo.Close()
		return
	}

	createdTime := time.Now().In(loc)
	updatedTime := time.Now().In(loc)
	place := &models.InsertPlace{
		Name:        "place",
		Description: "place description",
		Coordinates: models.Coordinates{
			Latitude:  "7.0591946",
			Longitude: "-73.8748526",
		},
		CreatedAt: createdTime,
		UpdatedAt: updatedTime,
	}

	result, err := repo.InsertPlace(context.Background(), place)
	if err != nil {
		t.Errorf("error insert place: %v", err)
		repo.Close()
		return
	}

	expected := &models.Place{
		ID:          result.ID,
		Name:        "place",
		Description: "place description",
		Coordinates: models.Coordinates{
			Latitude:  "7.0591946",
			Longitude: "-73.8748526",
		},
		CreatedAt: createdTime,
		UpdatedAt: updatedTime,
	}

	if result.Name != expected.Name {
		t.Errorf("Expected %s, but got %s", expected, result)
		repo.Close()
		return
	}

	repo.Close()
}

func TestUpdatePlace(t *testing.T) {
	//mogo conection
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	repo, err := database.NewMongoRepo(os.Getenv("DB_URI_TEST"))
	if err != nil {
		repo.Close()
		log.Fatal(err)
	}

	//code block to generate data
	ctx := context.Background()
	places, err := repo.ListPlaces(ctx)
	if err != nil {
		t.Errorf("Error searching places")
		repo.Close()
		return
	}
	if len(places) == 0 {
		t.Errorf("No places, create a new one")
		repo.Close()
		return
	}

	//location
	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		t.Errorf("error loading location: %v", err)
		repo.Close()
		return
	}
	createdTime := time.Now().In(loc)
	updatedTime := time.Now().In(loc)

	//data recipe
	place := &models.UpdatePlace{
		Name:        "place updated",
		Description: "place updated description",
		Coordinates: models.Coordinates{
			Latitude:  "7.0591946",
			Longitude: "-73.8748526",
		},
		UpdatedAt: updatedTime,
	}

	//test func
	result, err := repo.UpdatePlace(ctx, place, places[0].ID.Hex())
	if err != nil {
		t.Errorf("error insert place: %v", err)
		repo.Close()
		return
	}

	expected := &models.Place{
		ID:          result.ID,
		Name:        "place updated",
		Description: "place updated description",
		Coordinates: models.Coordinates{
			Latitude:  "7.0591946",
			Longitude: "-73.8748526",
		},
		CreatedAt: createdTime,
		UpdatedAt: updatedTime,
	}

	//print result
	if result.Name != expected.Name {
		t.Errorf("Expected %s, but got %s", expected, result)
		repo.Close()
		return
	}
	repo.Close()
}
