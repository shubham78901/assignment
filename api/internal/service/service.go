package service

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type CountryService interface {
	GetCountry(ctx context.Context, name string) (*model.Country, error)
}

type countryService struct {
	cache cache.CacheInterface
}

func NewCountryService(c cache.CacheInterface) CountryService {
	return &countryService{cache: c}
}

func (s *countryService) GetCountry(ctx context.Context, name string) (*model.Country, error) {
	log := logger.GetLogger()

	log.Info("Checking cache for country", zap.String("name", name))
	if country, found := s.cache.Get(ctx, name); found {
		log.Info("Country found in cache", zap.String("name", name))
		return &country, nil
	}

	log.Info("Fetching country data", zap.String("name", name))
	country, err := s.fetchCountryData(ctx, name)
	if err != nil {
		log.Error("Failed to fetch country data", zap.String("name", name), zap.Error(err))
		return nil, err
	}

	s.cache.Set(ctx, name, *country)
	log.Info("Stored country in cache", zap.String("name", name))
	return country, nil
}

func (s *countryService) fetchCountryData(ctx context.Context, name string) (*model.Country, error) {
	log := logger.GetLogger()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://restcountries.com/v3.1/name/%s", name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Error("Failed to create request", zap.String("url", url), zap.Error(err))
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("HTTP request failed", zap.String("url", url), zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Received non-OK response", zap.String("url", url), zap.Int("status", resp.StatusCode))
		return nil, fmt.Errorf("failed to fetch country data")
	}

	var countries []model.Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		log.Error("Failed to decode response", zap.String("url", url), zap.Error(err))
		return nil, err
	}

	if len(countries) == 0 {
		log.Warn("No country data found", zap.String("name", name))
		return nil, fmt.Errorf("no country found")
	}

	log.Info("Successfully fetched country data", zap.String("name", name))
	return &countries[0], nil
}
