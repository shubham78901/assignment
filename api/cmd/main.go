package main

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/router"
	"assignment/api/internal/service"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	log := logger.GetLogger()

	// Initialize dependencies
	c := cache.NewCache()
	svc := service.NewCountryService(c)
	r := router.SetupRouter(svc)

	// Configure HTTP server
	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		log.Info("Server starting on port 8000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error", zap.Error(err))
		}
	}()

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Wait for a termination signal
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server shutdown failed", zap.Error(err))
	} else {
		log.Info("Server gracefully stopped")
	}
}
