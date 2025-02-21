package handler

import (
	"assignment/api/internal/logger"
	"assignment/api/internal/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SearchCountryHandler handles searching for a country by name.
//
// @Summary Search country by name
// @Description Fetches details of a country using its name.
// @Tags Countries
// @Accept  json
// @Produce  json
// @Param name query string true "Country Name"
// @Success 200 {object} model.Country "Successfully retrieved country details"
// @Failure 400 {object} map[string]string "Bad Request - Missing country name"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/countries/search [get]
func SearchCountryHandler(svc service.CountryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetLogger()
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		name := c.Query("name")
		if name == "" {
			log.Warn("Missing country name in request")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing country name"})
			return
		}

		country, err := svc.GetCountry(ctx, name)
		if err != nil {
			log.Error("Failed to get country data", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Info("Country data retrieved successfully", zap.String("country", name))
		c.JSON(http.StatusOK, country)
	}
}
