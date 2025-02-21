package service

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/model"
	"context"
	"testing"

	"go.uber.org/zap"
)

func TestGetCountryFromCache(t *testing.T) {
	// Initialize logger
	logger.InitLogger()
	log := logger.GetLogger()

	c := cache.NewCache()
	svc := NewCountryService(c)
	ctx := context.Background()

	log.Info("Starting Test: TestGetCountryFromCache")

	// Set value in cache
	c.Set(ctx, "India", model.Country{Name: model.Name{Common: "India"}})
	log.Info("Set India in cache")

	// Fetch value from cache
	country, err := svc.GetCountry(ctx, "India")
	if err != nil {
		log.Error("Unexpected error", zap.Error(err))
		t.Errorf("Unexpected error: %v", err)
	}

	if country.Name.Common != "India" {
		log.Error("Expected India, got different value", zap.String("received", country.Name.Common))
		t.Errorf("Expected India, got %s", country.Name.Common)
	}

	log.Info("TestGetCountryFromCache passed")
}
