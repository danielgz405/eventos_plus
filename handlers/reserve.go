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
	"github.com/danielgz405/whale_places/utils"
	"github.com/gorilla/mux"
)

func InsertReserveHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = structures.InsertReserveRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}

		reserve := models.InsertReserve{
			UserID:  user.Id.Hex(),
			EventId: req.EventId,
			Acceted: false,
		}

		createdReserve, err := repository.InsertReserve(r.Context(), &reserve)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdReserve)
	}
}
func ListReservesByUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		params := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		reserves, err := repository.ListReservesByUser(r.Context(), params["user_id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		if reserves == nil {
			reserves = []models.Reserve{}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(reserves)
	}
}

func UpdateReserveHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var req = structures.UpdateReserveRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}

		layout := "Mon Jan 02 2006"
		// dates should have the same format "Fri Apr 05 2024"
		date, err := time.Parse(layout, req.DateToAcceted)
		if err != nil {
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}

		reserve := models.UpdateReserve{
			Acceted:       req.Acceted,
			DateToAcceted: date,
		}

		updatedReserve, err := repository.UpdateReserve(r.Context(), &reserve, params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedReserve)
	}
}

func GetReserveByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		reserve, err := repository.GetReserveById(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(reserve)
	}
}
