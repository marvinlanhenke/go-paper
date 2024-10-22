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
	"github.com/marvinlanhenke/go-paper/internal/repository"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	logger     *zap.SugaredLogger
	router     *chi.Mux
	server     *http.Server
	repository *repository.Repository
	config     *Config
}

func NewApplication(logger *zap.SugaredLogger, config *Config) (*Application, error) {
	app := &Application{
		logger: logger,
		router: chi.NewRouter(),
		config: config,
	}

	if err := app.setupDB(); err != nil {
		return nil, err
	}

	app.router.Use(middleware.RequestID)
	app.router.Use(middleware.RealIP)
	app.router.Use(middleware.Logger)
	app.router.Use(middleware.Recoverer)
	app.router.Use(middleware.Timeout(time.Second * 60))

	app.registerRoutes()

	return app, nil
}

func (app *Application) setupDB() error {
	db, err := gorm.Open(postgres.Open(app.config.db.addr), &gorm.Config{})
	if err != nil {
		app.logger.Errorw("error initializing database", "error", err)
		return err
	}
	if err := db.AutoMigrate(&repository.Paper{}); err != nil {
		app.logger.Errorw("error migrating database", "error", err)
		return err
	}

	app.repository = repository.New(db)

	return nil
}

func (app *Application) registerRoutes() {
	env := app.config.env
	version := app.config.version

	healthCheckHandler := handler.NewHealthCheckHandler(env, version)
	paperHandler := handler.NewPaperHandler(app.logger, app.repository)

	app.router.Route("/v1", func(r chi.Router) {
		r.Get("/health", healthCheckHandler.ServeHTTP)

		r.Route("/papers", func(r chi.Router) {
			r.Post("/", paperHandler.Create)
			r.Get("/", paperHandler.ReadAll)

			r.Route("/{id}", func(r chi.Router) {
				r.Use(paperHandler.WithPaperContext)

				r.Get("/", paperHandler.Read)
				r.Patch("/", paperHandler.Update)
				r.Delete("/", paperHandler.Delete)
			})
		})
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
