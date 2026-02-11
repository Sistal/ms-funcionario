package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort     string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	AllowedOrigins string
	AuthServiceURL string
}

func LoadConfig() *Config {
	// Cargar archivo .env si existe
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables or defaults")
	} else {
		log.Println("Loaded configuration from .env file")
	}

	config := &Config{
		ServerPort:     getEnv("SERVER_PORT", "8081"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "postgres"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:5174"),
		AuthServiceURL: getEnv("AUTH_SERVICE_URL", "http://localhost:8080"),
	}

	// Log de configuración (sin mostrar password)
	log.Printf("Database Config: Host=%s, Port=%s, User=%s, DBName=%s, SSLMode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBSSLMode)

	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Environment variable %s not set, using default: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
