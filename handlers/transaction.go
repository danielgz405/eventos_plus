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

func InsertTransactionHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = structures.InsertTransactionRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}

		layout := "Mon Jan 02 2006"
		// dates should have the same format "Fri Apr 05 2024"
		date, err := time.Parse(layout, req.CreateAt)
		if err != nil {
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}

		reserve := models.InsertTransaction{
			RecipientId: req.RecipientId,
			EmmiterId:   req.EmmiterId,
			EventId:     req.EventId,
			ReserveId:   req.RecipientId,
			CreateAt:    date,
			Total:       req.Total,
			Ref:         req.Ref,
			Subtotal:    req.Subtotal,
		}

		createdTransaction, err := repository.InsertTransaction(r.Context(), &reserve)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdTransaction)
	}
}
func ListTransactionsByUserHandler(s server.Server) http.HandlerFunc {
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
		reserves, err := repository.ListTransactionsByUser(r.Context(), params["user_id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		if reserves == nil {
			reserves = []models.Transaction{}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(reserves)
	}
}

func GetTransactionByIdHandler(s server.Server) http.HandlerFunc {
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
		reserve, err := repository.GetTransactionById(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(reserve)
	}
}
