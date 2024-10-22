package app

type Config struct {
	addr string
}

func NewConfig(addr string) *Config {
	return &Config{addr: addr}
}
