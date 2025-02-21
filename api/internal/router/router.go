package router

import (
	"assignment/api/internal/handler"
	"assignment/api/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns a Gin router with defined routes.
func SetupRouter(svc service.CountryService) *gin.Engine {
	r := gin.Default()

	// Group API routes under `/api`
	api := r.Group("/api")
	{
		countries := api.Group("/countries")
		{
			countries.GET("/search", handler.SearchCountryHandler(svc))
		}
	}

	return r
}
