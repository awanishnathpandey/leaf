package db

import (
	"context"

	"github.com/awanishnathpandey/leaf/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

// ConnectDB initializes the database connection pool
func ConnectDB() (*pgxpool.Pool, error) {
	// Load environment variables
	err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	// Get DB URL and max connections from the config
	dbURL := config.GetDatabaseURL()
	dbMaxConn := config.GetDBMaxConnections()

	// Initialize database connection pool
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to the database")
		return nil, err
	}

	// Set connection pool max count
	dbPool.Config().MaxConns = int32(dbMaxConn)

	return dbPool, nil
}
