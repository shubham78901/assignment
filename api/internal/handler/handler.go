package handler

import (
	"assignment/api/internal/logger" // Import model for Swagger response
	"assignment/api/internal/service"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// SearchCountryHandler handles searching for a country by name.
//
// @Summary Search country by name
// @Description Fetches details of a country using its name.
// @Tags Countries
// @Accept  json
// @Produce  json
// @Param name query string true "Country Name"
// @Success 200 {object} model.Country "Successfully retrieved country details"
// @Failure 400 {object} map[string]string "Bad Request - Missing country name"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /countries/search [get]
func SearchCountryHandler(svc service.CountryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		name := r.URL.Query().Get("name")
		if name == "" {
			log.Warn("Missing country name in request")
			http.Error(w, "Missing country name", http.StatusBadRequest)
			return
		}

		country, err := svc.GetCountry(ctx, name)
		if err != nil {
			log.Error("Failed to get country data", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Info("Country data retrieved successfully", zap.String("country", name))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(country)
	}
}
