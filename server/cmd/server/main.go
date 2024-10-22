package main

import (
	"github.com/marvinlanhenke/go-paper/internal/app"
	"go.uber.org/zap"
)

const addr = ":8080"

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Errorw("failed to sync logger", "error", err)
		}
	}()

	config := app.NewConfig(addr)

	app := app.NewApplication(logger, config)

	app.Run()
}
