package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config holds the database configuration values
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
}

func ReadEnv(path string) *Config {
	if path == "" {
		path = "./.env"
	}
	// Load .env file
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read and parse environment variables
	config := Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBName:     getEnv("DB_NAME", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
	}

	// Print the configuration (or you can use it as needed)
	//fmt.Printf("Config: %+v\n", config)
	return &config
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
