package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danielgz405/whale_places/middleware"
	"github.com/danielgz405/whale_places/models"
	"github.com/danielgz405/whale_places/repository"
	"github.com/danielgz405/whale_places/responses"
	"github.com/danielgz405/whale_places/server"
	"github.com/danielgz405/whale_places/structures"
	"github.com/gorilla/mux"
)

func InsertTypeEventHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = structures.InsertTypeEventRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}

		typeEvent := models.InsertTypeEvent{
			UserID:      user.Id.Hex(),
			Name:        req.Name,
			Description: req.Description,
		}

		createdTypeEvent, err := repository.InsertTypeEvent(r.Context(), &typeEvent)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdTypeEvent)
	}
}
func ListTypeEventsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		typeEvents, err := repository.ListTypeEvents(r.Context())
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		if typeEvents == nil {
			typeEvents = []models.TypeEvent{}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(typeEvents)
	}
}

func UpdateTypeEventHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var req = structures.UpdateTypeEventRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}

		typeEvent := models.UpdateTypeEvent{
			Name:        req.Name,
			Description: req.Description,
		}

		updatedTypeEvent, err := repository.UpdateTypeEvent(r.Context(), &typeEvent, params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedTypeEvent)
	}
}
func DeleteTypeEventHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err = repository.DeleteTypeEvent(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		responses.DeleteResponse(w, "TypeEvent deleted")
	}
}

func GetTypeEventByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		typeEvent, err := repository.GetTypeEventById(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(typeEvent)
	}
}
