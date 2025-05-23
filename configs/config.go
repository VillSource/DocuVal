package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Port        string
	Environment string
}

var AppConfig *Config

func GetAppConfig() *Config {
    if AppConfig == nil {
        LoadConfig()
    }
    return AppConfig
}

func LoadConfig() {
	_ = godotenv.Load()
	AppConfig = &Config{
		AppName:     "Docuval",
		Port:        getEnv("DOCUVAL_PORT", "3000"),
		Environment: getEnv("DOCUVAL_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
