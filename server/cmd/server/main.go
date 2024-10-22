package main

import (
	"github.com/joho/godotenv"
	_ "github.com/marvinlanhenke/go-paper/docs"
	"github.com/marvinlanhenke/go-paper/internal/app"
	"go.uber.org/zap"
)

// @title Go Paper API
// @version 0.0.1
// @description A simple API for managing your papers.
// @host localhost:8080
// @BasePath /v1
func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()

	if err := godotenv.Load(); err != nil {
		logger.Warnw("error while loading .env", "error", err)
	}

	app, err := app.NewApplication(logger, app.NewConfig())
	if err != nil {
		logger.Fatalw("error while initializing application", "error", err)
	}

	app.Run()
}
