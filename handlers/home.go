package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danielgz405/whale_places/server"
	"github.com/danielgz405/whale_places/utils"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	utils.DatabaseConnection(s)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to Prisma Home",
			Status:  http.StatusOK,
		})
	}
}
