package handler

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/model"
	"assignment/api/internal/service"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func TestSearchCountryHandler(t *testing.T) {
	// Initialize the logger before running the test
	logger.InitLogger()
	log := logger.GetLogger()

	// Initialize a valid cache instance
	c := cache.NewCache()
	svc := service.NewCountryService(c)

	// Set mock data in cache for predictable test results
	ctx := context.Background()
	mockCountry := model.Country{Name: model.Name{Common: "India"}}
	c.Set(ctx, "India", mockCountry)

	// Create a new HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", "/api/countries/search?name=India", nil)
	if err != nil {
		log.Error("Test setup failed: Error creating request", zap.Error(err))
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := SearchCountryHandler(svc)
	handler.ServeHTTP(rr, req)

	// Validate response status code
	if rr.Code != http.StatusOK {
		log.Error("Test failed: Expected status OK", zap.Int("status", rr.Code))
		t.Errorf("Expected status OK, got %v", rr.Code)
		return
	}

	log.Info("Test passed: Successfully received OK response")
}
