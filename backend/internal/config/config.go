package config

import "os"

type Config struct {
	Port        string
	GinMode     string
	DatabaseURL string
	RedisURL    string
}

func Load() (*Config, error) {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		GinMode:     getEnv("GIN_MODE", "debug"),
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/app_dev"),
		RedisURL:    getEnv("REDIS_URL", "redis://:redis@localhost:6379"),
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
