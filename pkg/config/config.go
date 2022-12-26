package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TeleToken   string
	PostgresURL string
	Debug       bool
	LogPath     string
	LogLevel    string
	UrlLink     string
	Webhook     bool
}

// GetConf returns a new Config struct
func GetConf() (*Config, error) {
	// Store the PATH environment variable in a variable
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	cfg := &Config{
		TeleToken:   getVarEnv("TELEGRAM_TOKEN", ""),
		PostgresURL: getUrlPostgres(),
		Debug:       boolEnv(getVarEnv("DEBUG", "true")),
		LogPath:     getVarEnv("LOG_PATH", "./log.txt"),
		LogLevel:    getVarEnv("LOG_LEVEL", "info"),
		UrlLink:     getVarEnv("URL_LINK", ""), // linkServ address
		Webhook:     boolEnv(getVarEnv("WEBHOOK_BOOL", "true")),
	}
	return cfg, nil
}

// Simple helper function to read an environment variable or return a default value
func getVarEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		log.Info(key + " = " + value)
		return value
	}
	log.Info("Default " + key + " = " + defaultVal)
	return defaultVal
}

func boolEnv(valEnv string) bool {
	if valEnv == "true" {
		return true
	} else {
		return false
	}
}

func getUrlPostgres() string {
	//db := getVarEnv("POSTGRES_DB", "")
	userDb := getVarEnv("POSTGRES_USER", "")
	passDb := getVarEnv("POSTGRES_PASSWORD", "")

	dbURL := fmt.Sprintf("postgres://%s:%s@postgres:5432/appdb", userDb, passDb)
	log.Info("URL Postgres - " + dbURL)
	return dbURL
}
