package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config consts: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
