package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"App"`
		Http `yaml:"Http"`
		Log  `yaml:"Log"`
		PG   `yaml:"PG"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	Http struct {
		Port string `yaml:"port"`
	}

	Log struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	}

	PG struct {
		DSN string `yaml:"DSN"`
		URL string `yaml:"URL"`
		//PoolMax int    `yaml:"PoolMax"` //todo разобраться с этим
	}
)

func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)

	if err != nil {
		return nil, fmt.Errorf("error of loading configuration: %w", err)
	}

	return cfg, nil
}
