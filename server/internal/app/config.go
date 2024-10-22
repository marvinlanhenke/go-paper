package app

import "github.com/marvinlanhenke/go-paper/internal/utils"

type Config struct {
	addr    string
	env     string
	version string
}

func NewConfig() *Config {
	addr := utils.GetString("ADDR", ":8080")
	env := utils.GetString("ENV", "development")
	version := utils.GetString("VERSION", "0.0.1")

	return &Config{
		addr:    addr,
		env:     env,
		version: version,
	}
}
