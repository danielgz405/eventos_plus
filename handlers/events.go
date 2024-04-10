package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/danielgz405/whale_places/middleware"
	"github.com/danielgz405/whale_places/models"
	"github.com/danielgz405/whale_places/repository"
	"github.com/danielgz405/whale_places/responses"
	"github.com/danielgz405/whale_places/server"
	"github.com/danielgz405/whale_places/structures"
	"github.com/danielgz405/whale_places/utils"
	"github.com/gorilla/mux"
)

func CreateEventHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Token validation
		profile, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}

		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := structures.InsertEventsRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println(err)
			responses.BadRequest(w, "Invalid request body")
			return
		}

		layout := "Mon Jan 02 2006"
		// dates should have the same format "Fri Apr 05 2024"
		date, err := time.Parse(layout, req.Date)
		if err != nil {
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}

		createEvent := models.InsertEvent{
			PlaceID:     req.PlaceID,
			TypeEvent:   req.TypeEvent,
			UserID:      profile.Id.Hex(),
			Name:        req.Name,
			Description: req.Description,
			Date:        date,
			IsFree:      req.IsFree,
			Capacity:    req.Capacity,
			TypeTiket:   req.TypeTiket,
		}

		events, err := repository.InsertEvent(r.Context(), &createEvent)
		if err != nil {
			responses.BadRequest(w, "Error creating events")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}
}

func GetEventByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		events, err := repository.GetEventById(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error getting events")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}
}

func UpdateEventHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		req := structures.UpdateEventsRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}

		updateEvent := models.UpdateEvent{
			TypeEvent: req.TypeEvent,
			Name:      req.Name,
			IsFree:    req.IsFree,
			Capacity:  req.Capacity,
			TypeTiket: req.TypeTiket,
		}
		events, err := repository.UpdateEvent(r.Context(), &updateEvent, params["id"])
		if err != nil {
			responses.BadRequest(w, "Error updating events")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}
}

func DeleteEventHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err = repository.DeleteEvent(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting events")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ListEventsByPageHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		limit, err := strconv.Atoi(params["limit"])
		if err != nil {
			responses.BadRequest(w, "Bad request")
			return
		}
		page, err := strconv.Atoi(params["page"])
		if err != nil {
			responses.BadRequest(w, "Bad request")
			return
		}
		events, quantity, err := repository.ListEventsByPage(r.Context(), limit, page)
		if err != nil {
			responses.BadRequest(w, "Error getting eventss")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.EventsResponse{
			Event:    events,
			Quantity: quantity,
		})
	}
}

func ListEventsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		events, err := repository.ListEvents(r.Context())
		if err != nil {
			responses.BadRequest(w, "Error getting events")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}
}

func ListEventsByNameHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		limit, err := strconv.Atoi(params["limit"])
		if err != nil {
			responses.BadRequest(w, "Bad request")
			return
		}
		page, err := strconv.Atoi(params["page"])
		if err != nil {
			responses.BadRequest(w, "Bad request")
			return
		}

		events, quantity, err := repository.ListEventsByName(r.Context(), limit, page, params["name"])
		if err != nil {
			responses.BadRequest(w, "Error getting eventss")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.EventsResponse{
			Event:    events,
			Quantity: quantity,
		})
	}
}
