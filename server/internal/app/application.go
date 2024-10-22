package app

import (
	"net/http"
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

	app.registerRoutes()

	return app
}

func (app *Application) registerRoutes() {
	healthCheckHandler := handler.NewHealthCheckHandler(app.logger)

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

	app.logger.Infow("starting server", "addr", app.config.addr)

	if err := app.server.ListenAndServe(); err != nil {
		app.logger.Fatal(err)
	}
}
