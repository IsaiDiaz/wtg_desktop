package config

import (
	"os"
)

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func GetServerPort() string {
	return GetEnv("SERVER_PORT", "8080")
}
