package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

// GetCacheExpiry retrieves the cache expiration duration from an environment variable.
func GetCacheExpiry() time.Duration {
	cacheExpiry := 5 * time.Minute // Default value
	// fmt.Println("CACHE_EXPIRY:", os.Getenv("CACHE_EXPIRY"))
	if cacheExpiryStr := os.Getenv("CACHE_EXPIRY"); cacheExpiryStr != "" {
		parsedValue, err := time.ParseDuration(cacheExpiryStr)
		if err != nil {
			log.Warn().Err(err).Msgf("Invalid value for CACHE_EXPIRY, using default: %v", cacheExpiry)
		} else {
			cacheExpiry = parsedValue
		}
	}
	return cacheExpiry
}

// GetCacheMaxSize returns the maximum size of the cache from the environment variable
func GetCacheMaxSize() int {
	cacheMaxSize := 100
	// fmt.Println("CACHE_MAX_SIZE:", os.Getenv("CACHE_MAX_SIZE"))
	if maxSizeStr := os.Getenv("CACHE_MAX_SIZE"); maxSizeStr != "" {
		var err error
		cacheMaxSize, err = strconv.Atoi(maxSizeStr)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid CACHE_MAX_SIZE value")
		}
	}
	return cacheMaxSize
}

// GetCacheCleanupInterval retrieves the cache cleanup interval duration from an environment variable.
func GetCacheCleanupInterval() time.Duration {
	cacheCleanupInterval := 10 * time.Minute // Default value
	// fmt.Println("CACHE_EXPIRY_INTERVAL:", os.Getenv("CACHE_EXPIRY_INTERVAL"))
	if cacheCleanupIntervalStr := os.Getenv("CACHE_EXPIRY_INTERVAL"); cacheCleanupIntervalStr != "" {
		parsedValue, err := time.ParseDuration(cacheCleanupIntervalStr)
		if err != nil {
			log.Warn().Err(err).Msgf("Invalid value for CACHE_EXPIRY_INTERVAL, using default: %v", cacheCleanupInterval)
		} else {
			cacheCleanupInterval = parsedValue
		}
	}
	return cacheCleanupInterval
}

// GetWorkerPoolSize returns the maximum worker pool size of the cache from the environment variable
func GetWorkerPoolSize() int {
	poolSize := 10
	// fmt.Println("WORKER_POOL_SIZE:", os.Getenv("WORKER_POOL_SIZE"))
	if maxSizeStr := os.Getenv("WORKER_POOL_SIZE"); maxSizeStr != "" {
		var err error
		poolSize, err = strconv.Atoi(maxSizeStr)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid WORKER_POOL_SIZE value")
		}
	}
	return poolSize
}

// GetWorkerPoolSize returns the maximum worker pool size of the cache from the environment variable
func GetAuthLastSeenQueueSize() int {
	queueSize := 100
	// fmt.Println("AUTH_LAST_SEEN_QUEUE_SIZE:", os.Getenv("AUTH_LAST_SEEN_QUEUE_SIZE"))
	if maxSizeStr := os.Getenv("AUTH_LAST_SEEN_QUEUE_SIZE"); maxSizeStr != "" {
		var err error
		queueSize, err = strconv.Atoi(maxSizeStr)
		if err != nil {
			log.Fatal().Err(err).Msg("Invalid AUTH_LAST_SEEN_QUEUE_SIZE value")
		}
	}
	return queueSize
}
