package repository

import (
	"fmt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/shop/internal/config"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

func NewPostgres(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDBName)
	logrus.Info(dsn)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to connect to postgres: %w", err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	return db, nil
}
