package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// LoadEnv loads the environment variables from the .env file
func LoadEnv() error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
		return err
	}
	return nil
}

// GetDatabaseURL returns the DATABASE_URL from the environment variables
func GetDatabaseURL() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbSSLMode == "" {
		log.Fatal().Msg("One or more required database environment variables (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME) are not set in .env file")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	return dbURL
}

// GetDBMaxConnections returns the max number of database connections from the environment
func GetDBMaxConnections() int {
	dbMaxConn := 20
	if maxConnStr := os.Getenv("DB_MAX_CONNECTIONS"); maxConnStr != "" {
		var err error
		dbMaxConn, err = strconv.Atoi(maxConnStr)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid DB_MAX_CONNECTIONS value")
		}
	}
	return dbMaxConn
}
