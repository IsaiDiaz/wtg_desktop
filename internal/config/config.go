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

func GetEnvironment() string {
	return GetEnv("APP_ENV", "development")
}

func GetJwtSecret() string {
	return GetEnv("JWT_SECRET", "")
}

func GetJwtRefreshSecret() string {
	return GetEnv("JWT_REFRESH_SECRET", "")
}
