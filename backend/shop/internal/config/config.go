package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost     string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	ServerPort     string `env:"SERVER_PORT" env-required:"true"`
	PostgresUser   string `env:"POSTGRES_USER" required:"true"`
	PostgresPass   string `env:"POSTGRES_PASSWORD" required:"true"`
	PostgresHost   string `env:"POSTGRES_HOST" required:"true"`
	PostgresPort   string `env:"POSTGRES_PORT" required:"true"`
	PostgresDBName string `env:"POSTGRES_DB_NAME" required:"true"`
	PathPublicKey  string `env:"PATH_PUBLIC_KEY" required:"true"`
	AppEnv         string `env:"APP_ENV" envDefault:"dev"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	godotenv.Load(".env")
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	return &cfg, nil
}
