package app

import "github.com/marvinlanhenke/go-paper/internal/utils"

type Config struct {
	addr string
}

func NewConfig() *Config {
	addr := utils.GetString("ADDR", ":8080")

	return &Config{addr: addr}
}
