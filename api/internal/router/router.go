// internal/router/router.go
package router

import (
	"assignment/api/internal/handler"
	"assignment/api/internal/service"
	"net/http"
)

func SetupRouter(svc service.CountryService) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/api/countries/search", handler.SearchCountryHandler(svc))
	return r
}
