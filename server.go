package main

import (
	"os"
	"os/signal"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/internal/database"
	"github.com/awanishnathpandey/leaf/internal/middleware"
	"github.com/awanishnathpandey/leaf/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// Set zerolog to output JSON format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(os.Stdout)

	// Initialize database connection
	dbPool, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Error initializing database connection")
		return
	}
	// Defer closing the database connection pool
	defer dbPool.Close()

	queries := generated.New(dbPool)
	// Initialize middleware with queries object
	middleware.InitializeQueries(queries)
	middleware.StartWorkerPool()
	middleware.StartAuditWorkerPool()
	middleware.InitializePermissionCache()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{})

	// Setup routes using the routes package
	routes.SetupRoutes(app, queries, dbPool)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info().Msg("Gracefully shutting down...")
		// Stop worker pool
		middleware.StopWorkerPool()
		middleware.StopAuditWorkerPool()
		middleware.StopPermissionCacheCleanup()
		_ = app.Shutdown()
	}()

	// Start server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal().Err(err).Msg("Error starting server")
	}
	log.Info().Msg("Server stopped")
}
