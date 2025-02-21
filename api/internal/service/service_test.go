package service

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/model"
	"context"
	"testing"
)

func TestGetCountryFromCache(t *testing.T) {
	c := cache.NewCache()
	svc := NewCountryService(c)
	ctx := context.Background()

	// Correctly initializing model.Name struct
	c.Set(ctx, "India", model.Country{Name: model.Name{Common: "India"}})

	country, err := svc.GetCountry(ctx, "India")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if country.Name.Common != "India" {
		t.Errorf("Expected India, got %s", country.Name.Common)
	}
}
