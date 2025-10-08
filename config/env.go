package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL        string
	ServerPort         string
	JWTSecretKey       string
	JWTExpirationHours time.Duration
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("File .env tidak ditemukan, membaca dari environment variables")
	}

	dbUser := getEnv("DB_USER", "dsiunair")
	dbPassword := getEnv("DB_PASSWORD", "root")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "golang-train")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	serverPort := getEnv("SERVER_PORT", "4000")
	jwtSecret := getEnv("JWT_SECRET_KEY", "default_secret_key")
	jwtExpHoursStr := getEnv("JWT_EXPIRATION_HOURS", "72")

	jwtExpHours, err := strconv.Atoi(jwtExpHoursStr)
	if err != nil {
		return nil, fmt.Errorf("JWT_EXPIRATION_HOURS tidak valid: %w", err)
	}

	return &Config{
		DatabaseURL:        databaseURL,
		ServerPort:         serverPort,
		JWTSecretKey:       jwtSecret,
		JWTExpirationHours: time.Duration(jwtExpHours) * time.Hour,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
