package handler

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/service"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchCountryHandler(t *testing.T) {
	// Initialize the logger before running the test
	logger.InitLogger()
	log := logger.GetLogger()

	// Initialize a valid cache instance
	c := cache.NewCache()
	svc := service.NewCountryService(c)

	req, err := http.NewRequestWithContext(context.Background(), "GET", "/api/countries/search?name=India", nil)
	if err != nil {
		log.Error("Test setup failed: Error creating request", logger.ErrorField(err))
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := SearchCountryHandler(svc)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Error("Test failed: Expected status OK", logger.ErrorField(err))
		t.Errorf("Expected status OK, got %v", rr.Code)
	} else {
		log.Info("Test passed: Successfully received OK response")
	}
}
