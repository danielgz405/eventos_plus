package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danielgz405/whale_places/middleware"
	"github.com/danielgz405/whale_places/models"
	"github.com/danielgz405/whale_places/repository"
	"github.com/danielgz405/whale_places/responses"
	"github.com/danielgz405/whale_places/server"
	"github.com/danielgz405/whale_places/structures"
	"github.com/gorilla/mux"
)

func InsertPlaceHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = structures.InsertPlaceRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		place := models.InsertPlace{}
		loc, error := time.LoadLocation("America/Bogota")
		if error != nil {
			responses.InternalServerError(w, "Invalid request")
			return
		}

		place = models.InsertPlace{
			Name:        req.Name,
			Description: req.Description,
			Coordinates: models.Coordinates{
				Latitude:  req.Coordinates.Latitude,
				Longitude: req.Coordinates.Longitude,
			},
			CreatedAt: time.Now().In(loc),
			UpdatedAt: time.Now().In(loc),
		}

		createdPlace, err := repository.InsertPlace(r.Context(), &place)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		repository.AuditOperation(r.Context(), *user, "place", "created")
		if err != nil {
			responses.InternalServerError(w, "Audit error")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdPlace)
	}
}
func ListPlacesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		places, err := repository.ListPlaces(r.Context())
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		if places == nil {
			places = []models.Place{}
		}
		repository.AuditOperation(r.Context(), *user, "place", "read")
		if err != nil {
			responses.InternalServerError(w, "Audit error")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(places)
	}
}

func UpdatePlaceHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var req = structures.UpdatePlaceRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		loc, error := time.LoadLocation("America/Bogota")
		if error != nil {
			responses.InternalServerError(w, "Invalid request")
			return
		}
		place := models.UpdatePlace{}

		place = models.UpdatePlace{
			Name:        req.Name,
			Description: req.Description,
			Coordinates: models.Coordinates{
				Latitude:  req.Coordinates.Latitude,
				Longitude: req.Coordinates.Longitude,
			},
			UpdatedAt: time.Now().In(loc),
		}

		updatedPlace, err := repository.UpdatePlace(r.Context(), &place, params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		repository.AuditOperation(r.Context(), *user, "place", "updated")
		if err != nil {
			responses.InternalServerError(w, "Audit error")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedPlace)
	}
}
func DeletePlaceHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err = repository.DeletePlace(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		repository.AuditOperation(r.Context(), *user, "place", "delete")
		if err != nil {
			responses.InternalServerError(w, "Audit error")
			return
		}
		responses.DeleteResponse(w, "Place deleted")
	}
}

func GetPlaceByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		place, err := repository.GetPlaceById(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		repository.AuditOperation(r.Context(), *user, "place", "read")
		if err != nil {
			responses.InternalServerError(w, "Audit error")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(place)
	}
}
