package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURI         string
	MongoDBName        string
	ServerPort         string
	JWTSecretKey       string
	JWTExpirationHours time.Duration
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("File .env tidak ditemukan, membaca dari environment variables")
	}

	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	mongoDBName := getEnv("MONGODB_NAME", "golang-train")

	serverPort := getEnv("SERVER_PORT", "4000")
	jwtSecret := getEnv("JWT_SECRET_KEY", "ApalahR4has!a!N!")
	jwtExpHoursStr := getEnv("JWT_EXPIRATION_HOURS", "72")

	jwtExpHours, err := strconv.Atoi(jwtExpHoursStr)
	if err != nil {
		return nil, fmt.Errorf("JWT_EXPIRATION_HOURS tidak valid: %w", err)
	}

	return &Config{
		MongoDBURI:         mongoURI,
		MongoDBName:        mongoDBName,
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
