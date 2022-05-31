package main

import "github.com/caarlos0/env/v6"

type Config struct {
	TCPPort int32 `env:"TCP_PORT" envDefault:"8080"`
}

func parseConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
