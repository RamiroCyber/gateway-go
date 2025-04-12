package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Load() error
	Get(key string) string
}

type EnvConfig struct {
	loaded bool
}

func NewEnvConfig() *EnvConfig {
	return &EnvConfig{}
}

func (e *EnvConfig) Load() error {
	if !e.loaded {
		if err := godotenv.Load(); err != nil {
			return err
		}
		e.loaded = true
	}
	return nil
}

func (e *EnvConfig) Get(key string) string {
	return os.Getenv(key)
}
