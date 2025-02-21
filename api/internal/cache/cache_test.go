// internal/cache/cache_test.go
package cache

import (
	"assignment/api/internal/logger"
	"assignment/api/internal/model"
	"context"
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	// Initialize the logger before running the test
	logger.InitLogger()
	log := logger.GetLogger()

	c := NewCache()
	ctx := context.Background()

	// Correctly initializing model.Name struct
	country := model.Country{Name: model.Name{Common: "India"}}

	log.Info("Setting country in cache", logger.ErrorField(nil))
	c.Set(ctx, "India", country)

	result, found := c.Get(ctx, "India")

	if !found {
		log.Error("Test failed: Expected country to be found in cache")
		t.Errorf("Expected country to be found in cache")
	} else {
		log.Info("Country found in cache", logger.ErrorField(nil))
	}

	if result.Name.Common != "India" {
		log.Error("Test failed: Expected country name to be India", logger.ErrorField(nil))
		t.Errorf("Expected country name to be India, got %s", result.Name.Common)
	} else {
		log.Info("Test passed: Country name matched expected value")
	}
}
