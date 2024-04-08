package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/danielgz405/whale_places/handlers"
	"github.com/danielgz405/whale_places/middleware"
	"github.com/danielgz405/whale_places/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URI := os.Getenv("DB_URI")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      ":" + PORT,
		JWTSecret: JWT_SECRET,
		DbURI:     DB_URI,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/welcome", handlers.HomeHandler(s)).Methods(http.MethodGet)

	//Auth
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	// User routes
	r.HandleFunc("/profile", handlers.ProfileHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/user/update/{id}", handlers.UpdateUserHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/user/delete/{id}", handlers.DeleteUserHandler(s)).Methods(http.MethodDelete)

	// Place routes
	r.HandleFunc("/place", handlers.InsertPlaceHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/place/one/{id}", handlers.GetPlaceByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/place/list", handlers.ListPlacesHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/place/update/{id}", handlers.UpdatePlaceHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/place/delete/{id}", handlers.DeletePlaceHandler(s)).Methods(http.MethodDelete)

	// types events routes
	r.HandleFunc("/type_event", handlers.InsertTypeEventHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/type_event/one/{id}", handlers.GetTypeEventByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/type_event/list", handlers.ListTypeEventsHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/type_event/update/{id}", handlers.UpdateTypeEventHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/type_event/delete/{id}", handlers.DeleteTypeEventHandler(s)).Methods(http.MethodDelete)

	// reserve routes
	r.HandleFunc("/reserve", handlers.InsertReserveHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/reserve/one/{id}", handlers.GetReserveByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/reserve/list/{user_id}", handlers.ListReservesByUserHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/reserve/update/{id}", handlers.UpdateReserveHandler(s)).Methods(http.MethodPatch)
}
