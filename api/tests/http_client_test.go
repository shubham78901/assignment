// tests/http_client_test.go
package tests

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/service"
	"context"
	"testing"
)

func TestFetchCountryData(t *testing.T) {
	// Initialize the logger before running the test
	logger.InitLogger()
	log := logger.GetLogger()

	// Initialize a valid cache instance
	c := cache.NewCache()
	svc := service.NewCountryService(c)
	ctx := context.Background()

	_, err := svc.GetCountry(ctx, "India")
	if err != nil {
		log.Error("Test failed: Expected successful fetch", logger.ErrorField(err))
		t.Errorf("Expected successful fetch, got error: %v", err)
	} else {
		log.Info("Test passed: Successfully fetched country data")
	}
}
