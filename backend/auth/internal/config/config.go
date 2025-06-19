package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost     string `env:"SERVER_HOST" binding:"required"`
	ServerPort     string `env:"SERVER_PORT" binding:"required"`
	AppEnv         string `env:"APP_ENV" envDefault:"prod"`
	PathPrivateKey string `env:"PATH_PRIVATE_KEY" binding:"required"`
	PathPublicKey  string `env:"PATH_PUBLIC_KEY" binding:"required"`

	PostgresUser   string `env:"POSTGRES_USER" required:"true"`
	PostgresPass   string `env:"POSTGRES_PASSWORD" required:"true"`
	PostgresHost   string `env:"POSTGRES_HOST" required:"true"`
	PostgresPort   string `env:"POSTGRES_PORT" required:"true"`
	PostgresDBName string `env:"POSTGRES_DB_NAME" required:"true"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	godotenv.Load(".env")
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
