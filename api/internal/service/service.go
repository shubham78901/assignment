// internal/service/service.go
package service

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	if country, found := s.cache.Get(ctx, name); found {
		return &country, nil
	}

	country, err := s.fetchCountryData(ctx, name)
	if err != nil {
		return nil, err
	}

	s.cache.Set(ctx, name, *country)
	return country, nil
}

func (s *countryService) fetchCountryData(ctx context.Context, name string) (*model.Country, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://restcountries.com/v3.1/name/%s", name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch country data")
	}

	var countries []model.Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, err
	}

	if len(countries) == 0 {
		return nil, fmt.Errorf("no country found")
	}
	return &countries[0], nil
}
