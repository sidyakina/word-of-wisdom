package main

import (
	"errors"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	TCPPort        int32  `env:"TCP_PORT" envDefault:"8080"`
	QuotesFilePath string `env:"QUOTES_FILE_PATH"`
}

func parseConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	if cfg.QuotesFilePath == "" {
		return nil, errors.New("empty quotes file path")
	}

	return cfg, nil
}
