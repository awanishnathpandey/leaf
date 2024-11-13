package config

import (
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
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal().Msg("DATABASE_URL not set in .env file")
	}
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
