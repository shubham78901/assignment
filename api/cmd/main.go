package main

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/router"
	"assignment/api/internal/service"

	_ "assignment/api/docs" // Import generated Swagger docs

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Country API
// @version 1.0
// @description This is an API for fetching country information.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /
func main() {
	logger.InitLogger()
	log := logger.GetLogger()

	// Initialize dependencies
	c := cache.NewCache()
	svc := service.NewCountryService(c)
	r := router.SetupRouter(svc)

	// ✅ Add Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configure HTTP server
	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	// Start server
	go func() {
		log.Info("Server starting on port 8000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error", zap.Error(err))
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server shutdown failed", zap.Error(err))
	} else {
		log.Info("Server gracefully stopped")
	}
}
