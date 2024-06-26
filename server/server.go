package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	database "github.com/danielgz405/whale_places/database"
	repository "github.com/danielgz405/whale_places/repository"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port      string
	JWTSecret string
	DbURI     string
	DB_URI_2  string
	DB_URI_3  string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DbURI == "" {
		return nil, errors.New("database uri is required")
	}
	if config.DB_URI_2  == "" {
		return nil, errors.New("database uri is required 2")
	}
	if config.DB_URI_3  == "" {
		return nil, errors.New("database uri is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})

	handler := c.Handler(b.router)
	repo, err := database.NewMongoRepo(b.config.DbURI)
	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)
	log.Println("Server started on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
