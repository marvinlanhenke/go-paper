package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marvinlanhenke/go-paper/internal/handler"
	"go.uber.org/zap"
)

type Application struct {
	logger *zap.SugaredLogger
	router *chi.Mux
	server *http.Server
	config *Config
}

func NewApplication(logger *zap.SugaredLogger, config *Config) *Application {
	app := &Application{
		logger: logger,
		router: chi.NewRouter(),
		config: config,
	}

	app.router.Use(middleware.RequestID)
	app.router.Use(middleware.RealIP)
	app.router.Use(middleware.Logger)
	app.router.Use(middleware.Recoverer)
	app.router.Use(middleware.Timeout(time.Second * 60))

	app.registerRoutes()

	return app
}

func (app *Application) registerRoutes() {
	env := app.config.env
	version := app.config.version

	healthCheckHandler := handler.NewHealthCheckHandler(app.logger, env, version)

	app.router.Route("/v1", func(r chi.Router) {
		r.Get("/health", healthCheckHandler.ServeHTTP)
	})
}

func (app *Application) Run() {
	app.server = &http.Server{
		Addr:         app.config.addr,
		Handler:      app.router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := app.server.Shutdown(ctx); err != nil {
			app.logger.Errorw("error while shutting down server", "error", err)
		}
	}()

	app.logger.Infow("starting server", "addr", app.config.addr)

	err := app.server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		app.logger.Fatal(err)
	}

	app.logger.Infow("stopping server", "addr", app.config.addr)
}
