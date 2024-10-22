package main

import (
	"github.com/joho/godotenv"
	"github.com/marvinlanhenke/go-paper/internal/app"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()

	if err := godotenv.Load(); err != nil {
		logger.Warnw("error while loading .env", "error", err)
	}

	app := app.NewApplication(logger, app.NewConfig())

	app.Run()
}
