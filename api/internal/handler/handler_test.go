package handler

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/model"
	"assignment/api/internal/service"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	mockCountry := model.Country{
		Name:       model.Name{Common: "India"},
		Capital:    []string{"New Delhi"}, // Capital as a slice
		Currencies: map[string]model.Currency{"INR": {Name: "Indian Rupee", Symbol: "₹"}},
		Population: 1400000000,
	}
	c.Set(ctx, "India", mockCountry)

	// Create a new Gin router and register the handler
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/countries/search", SearchCountryHandler(svc))

	// Create a new HTTP request with the search query
	req, err := http.NewRequest(http.MethodGet, "/api/countries/search?name=India", nil)
	if err != nil {
		log.Error("Test setup failed: Error creating request", zap.Error(err))
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Validate response status code
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

	// Parse the response body
	var response model.Country
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error unmarshaling response: %v", err)
	}

	// Validate response data
	assert.Equal(t, mockCountry, response, "Expected country data to match the mock country")

	log.Info("Test passed: Successfully received expected country response")
}
