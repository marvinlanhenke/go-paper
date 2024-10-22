package app

import "github.com/marvinlanhenke/go-paper/internal/utils"

type Config struct {
	addr    string
	env     string
	version string
	db      *dbConfig
}

type dbConfig struct {
	addr string
}

func NewConfig() *Config {
	addr := utils.GetString("ADDR", ":8080")
	env := utils.GetString("ENV", "development")
	version := utils.GetString("VERSION", "0.0.1")

	dbConfig := &dbConfig{
		addr: utils.GetString("DB_ADDR", "postgres://admin:admin@localhost:5432/gopaper?sslmode=disable"),
	}

	return &Config{
		addr:    addr,
		env:     env,
		version: version,
		db:      dbConfig,
	}
}
