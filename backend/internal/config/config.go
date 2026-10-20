package config

import (
	"errors"

	"github.com/romina-kh/NetEng-Project/backend/internal/db"
)

type Config struct {
	Postgres *db.PostGresConfig
	JWT      JWTConfig
	Server   ServerConfig
}

type JWTConfig struct {
	Secret string
}

type ServerConfig struct {
	Address string
}

func LoadConfig() (*Config, error) {

	cfg := &Config{
		Postgres: &db.PostGresConfig{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnv("DB_PORT", "5432"),
			User:     GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", ""),
			DBName:   GetEnv("DB_NAME", "techyar"),
		},
		JWT: JWTConfig{
			Secret: GetEnv("JWT_SECRET", ""),
		},
		Server: ServerConfig{
			Address: GetEnv("SERVER_ADDRESS", ":8080"),
		},
	}

	if cfg.Postgres.Password == "" {
		return nil, errors.New("DB_PASSWORD is required")
	}

	if cfg.JWT.Secret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}

	return cfg, nil
}
