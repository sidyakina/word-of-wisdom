package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	TCPServerAddress string        `env:"TCP_SERVER_ADDRESS"`
	ReadTimeout      time.Duration `env:"READ_TIMEOUT" envDefault:"2s"`
	NumberQuotes     int           `env:"NUMBER_QUOTES" envDefault:"1"`
}

func parseConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	if cfg.TCPServerAddress == "" {
		return nil, fmt.Errorf("wrong TCP server address %v", cfg.TCPServerAddress)
	}

	if cfg.NumberQuotes < 1 {
		return nil, fmt.Errorf("wrong number quotes %v", cfg.NumberQuotes)
	}

	if cfg.ReadTimeout <= 0 {
		return nil, fmt.Errorf("wrong timeout %v", cfg.ReadTimeout)
	}

	return cfg, nil
}
