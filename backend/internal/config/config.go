package config

import "github.com/romina-kh/NetEng-Project/backend/internal/db"

type Config struct {
	Server   ServerConfig
	Postgres *db.PostGresConfig
}

type ServerConfig struct {
	Address string
}

func LoadConfig() (*Config, error) {

	cfg := &Config{
		Server: ServerConfig{
			Address: GetEnv("SERVER_ADDRESS", ":8080"),
		},
		Postgres: &db.PostGresConfig{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnv("DB_PORT", "5432"),
			User:     GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", "netengpass"),
			DBName:   GetEnv("DB_NAME", "techyar"),
		},
	}

	return cfg, nil
}