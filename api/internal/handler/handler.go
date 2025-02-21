// internal/handler/handler.go
package handler

import (
	"assignment/api/internal/logger"
	"assignment/api/internal/service"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/zap"
)

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
			log.Error("Failed to get country data", logger.ErrorField(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Info("Country data retrieved successfully", zap.String("country", name))
		json.NewEncoder(w).Encode(country)
	}
}
