package repository

import (
	"context"

	"github.com/danielgz405/whale_places/models"
)

type Repository interface {
	//Users
	InsertUser(ctx context.Context, user *models.InsertUser) (*models.Profile, error)
	GetUserById(ctx context.Context, id string) (*models.Profile, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, data models.UpdateUser) (*models.Profile, error)
	DeleteUser(ctx context.Context, id string) error

	//Place
	InsertPlace(ctx context.Context, place *models.InsertPlace) (*models.Place, error)
	UpdatePlace(ctx context.Context, data *models.UpdatePlace, id string) (*models.Place, error)
	GetPlaceById(ctx context.Context, id string) (*models.Place, error)
	ListPlaces(ctx context.Context) ([]models.Place, error)
	DeletePlace(ctx context.Context, id string) error

	//type events
	InsertTypeEvent(ctx context.Context, typeEvent *models.InsertTypeEvent) (*models.TypeEvent, error)
	UpdateTypeEvent(ctx context.Context, data *models.UpdateTypeEvent, id string) (*models.TypeEvent, error)
	GetTypeEventById(ctx context.Context, id string) (*models.TypeEvent, error)
	ListTypeEvents(ctx context.Context) ([]models.TypeEvent, error)
	DeleteTypeEvent(ctx context.Context, id string) error

	//type events
	InsertReserve(ctx context.Context, reserve *models.InsertReserve) (*models.Reserve, error)
	UpdateReserve(ctx context.Context, data *models.UpdateReserve, id string) (*models.Reserve, error)
	GetReserveById(ctx context.Context, id string) (*models.Reserve, error)
	ListReservesByUser(ctx context.Context, userId string) ([]models.Reserve, error)

	//audit
	AuditOperation(ctx context.Context, user models.Profile, table string, operationType string) error

	//Close the connection
	Close() error
}

var implementation Repository

// Repo
func SetRepository(repository Repository) {
	implementation = repository
}

// Close the connection
func Close() error {
	return implementation.Close()
}
