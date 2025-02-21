package router

import (
	"assignment/api/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(svc service.CountryService) *gin.Engine {
	r := gin.Default()
	// Define your API routes here
	return r
}
