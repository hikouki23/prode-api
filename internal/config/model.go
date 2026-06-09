package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port            int    `env:"PORT,required"`
	MaxOpenConns    int    `env:"MAX_OPEN_CONNS"`
	MaxIdleConns    int    `env:"MAX_IDLE_CONNS"`
	ConnMaxIdleTime int    `env:"CONN_MAX_IDLE_TIME"`
	ConnMaxLifetime int    `env:"CONN_MAX_LIFETIME"`
	DatabaseURL     string `env:"DATABASE_URL,required"`
}

func LoadConfig() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
