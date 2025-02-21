package tests

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/service"
	"context"
	"testing"

	"go.uber.org/zap"
)

func TestFetchCountryData(t *testing.T) {
	// Initialize the logger before running the test
	logger.InitLogger()
	log := logger.GetLogger()

	// Initialize a valid cache instance
	c := cache.NewCache()
	svc := service.NewCountryService(c)
	ctx := context.Background()

	log.Info("Starting TestFetchCountryData")

	// Test fetching country data
	_, err := svc.GetCountry(ctx, "India")
	if err != nil {
		log.Error("Test failed: Expected successful fetch", zap.Error(err))
		t.Errorf("Expected successful fetch, got error: %v", err)
	} else {
		log.Info("Test passed: Successfully fetched country data")
	}
}
