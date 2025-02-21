// cmd/main.go
package main

import (
	"assignment/api/internal/cache"
	"assignment/api/internal/logger"
	"assignment/api/internal/router"
	"assignment/api/internal/service"
	"context"
	"net/http"
	"time"
)

func main() {
	logger.InitLogger()
	log := logger.GetLogger()

	c := cache.NewCache()
	svc := service.NewCountryService(c)
	r := router.SetupRouter(svc)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error", logger.ErrorField(err))
		}
	}()

	log.Info("Server running on port 8000")

	shutdown := make(chan struct{})
	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
