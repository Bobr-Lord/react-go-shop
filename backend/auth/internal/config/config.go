package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost string `env:"SERVER_HOST" binding:"required"`
	ServerPort string `env:"SERVER_PORT" binding:"required"`
	AppEnv     string `env:"APP_ENV" envDefault:"prod"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	godotenv.Load(".env")
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
