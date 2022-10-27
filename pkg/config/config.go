package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TeleToken string
	Debug     bool
	LogPath   string
	LogLevel  string
}

// GetConf returns a new Config struct
func GetConf() (*Config, error) {
	// Store the PATH environment variable in a variable
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{
		TeleToken: getVarEnv("TELEGRAM_TOKEN", ""),
		Debug:     boolEnv(getVarEnv("DEBUG", "true")),
		LogPath:   getVarEnv("LOG_PATH", "./log.txt"),
		LogLevel:  getVarEnv("LOG_LEVEL", "info"),
	}
	return cfg, nil
}

// Simple helper function to read an environment variable or return a default value
func getVarEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func boolEnv(valEnv string) bool {
	if valEnv == "true" {
		return true
	} else {
		return false
	}
}
