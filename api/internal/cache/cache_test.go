package cache

import (
	"assignment/api/internal/logger"
	"assignment/api/internal/model"
	"context"
	"testing"

	"go.uber.org/zap"
)

func TestCacheSetAndGet(t *testing.T) {
	// Initialize the logger before running the test
	logger.InitLogger()
	log := logger.GetLogger()

	// Initialize the cache instance
	c := NewCache()
	ctx := context.Background()

	// Create a mock country entry
	country := model.Country{Name: model.Name{Common: "India"}}

	log.Info("Setting country in cache", zap.String("country", country.Name.Common))
	c.Set(ctx, "India", country)

	// Retrieve from cache
	result, found := c.Get(ctx, "India")

	if !found {
		log.Error("Test failed: Expected country to be found in cache")
		t.Fatalf("Expected country to be found in cache")
	}

	log.Info("Country found in cache", zap.String("country", result.Name.Common))

	// Validate country name
	if result.Name.Common != "India" {
		log.Error("Test failed: Country name does not match", zap.String("expected", "India"), zap.String("got", result.Name.Common))
		t.Errorf("Expected country name to be 'India', got '%s'", result.Name.Common)
	} else {
		log.Info("Test passed: Country name matched expected value")
	}
}
