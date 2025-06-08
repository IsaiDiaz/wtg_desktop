package config

import (
	"os"
	"strconv"
	"time"
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

func IsUDPDiscoveryEnabled() bool {
	return GetEnv("UDP_DISCOVERY_ENABLED", "false") == "true"
}

func GetUDPDiscoveryPort() int {
	portStr := GetEnv("UDP_DISCOVERY_PORT", "7777")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 7777
	}
	return port
}

func GetUDPDiscoveryBroadcastIP() string {
	return GetEnv("UDP_DISCOVERY_BROADCAST_IP", "255.255.255.255")
}

func GetUDPDiscoveryInterval() time.Duration {
	internalStr := GetEnv("UDP_DISCOVERY_INTERNAL", "5s")
	duration, err := time.ParseDuration(internalStr)
	if err != nil {
		return 5 * time.Second
	}
	return duration
}
