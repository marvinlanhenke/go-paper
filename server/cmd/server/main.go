package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marvinlanhenke/go-paper/internal/handler"
	"go.uber.org/zap"
)

const addr = ":8080"

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", handler.HealthCheckHandler)
	})

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Errorw("failed to sync logger", "error", err)
		}
	}()

	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
	}

	logger.Infow("starting server", "addr", addr)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
