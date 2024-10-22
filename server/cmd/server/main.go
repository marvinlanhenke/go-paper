package main

import (
	"github.com/marvinlanhenke/go-paper/internal/app"
	"go.uber.org/zap"
)

const addr = ":8080"

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()

	config := app.NewConfig(addr)

	app := app.NewApplication(logger, config)

	app.Run()
}
