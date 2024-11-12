package main

import (
	// "log"
	// "net/http"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/awanishnathpandey/leaf/graph"
	"github.com/awanishnathpandey/leaf/graph/resolvers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Set zerolog to output JSON format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(os.Stdout)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{})

	// Custom logger middleware for Fiber using zerolog
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next() // Process request

		// Log request details
		log.Info().
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("status", c.Response().StatusCode()).
			Dur("latency", time.Since(start)).
			Msg("Request processed")

		return err
	})

	// Initialize database connection
	//  dbConn, err := db.Connect()
	//  if err != nil {
	// 	 log.Fatal().Err(err).Msg("Could not connect to the database")
	//  }
	//  defer dbConn.Close()

	// GraphQL handler using gqlgen
	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))
	//  graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{DB: dbConn}}))

	// Set up GraphQL endpoint with Fiber-compatible adapter
	app.Post("/graphql", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		graphqlHandler.ServeHTTP(w, r)
	}))

	// Set up GraphQL Playground with Fiber-compatible adapter
	app.Get("/playground", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		playground.Handler("GraphQL Playground", "/graphql").ServeHTTP(w, r)
	}))

	// Simple health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		log.Info().Msg("Health check accessed")
		return c.SendString("OK")
	})

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info().Msg("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// Start server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal().Err(err).Msg("Error starting server")
	}

	//  log.Info().Msg("Running cleanup tasks...")
	//  if err := dbConn.Close(); err != nil {
	// 	 log.Error().Err(err).Msg("Error closing database connection")
	//  }
	log.Info().Msg("Server stopped")
}
