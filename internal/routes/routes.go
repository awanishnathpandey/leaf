package routes

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/graph"
	"github.com/awanishnathpandey/leaf/graph/resolvers"
	"github.com/awanishnathpandey/leaf/internal/middleware"
	gqlprometheus "github.com/awanishnathpandey/leaf/internal/prometheus"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/rs/zerolog/log"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(app *fiber.App, queries *generated.Queries) {

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

	// Add JWT authentication middleware to protect the GraphQL route
	app.Use("/graphql", middleware.JWTMiddleware(queries))

	// GraphQL handler using gqlgen
	// graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{DB: queries}}))

	// Set up GraphQL endpoint with Fiber-compatible adapter
	// app.Post("/graphql", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	graphqlHandler.ServeHTTP(w, r)
	// }))

	// GraphQL handler using gqlgen
	graphqlHandler := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{DB: queries}}))

	// Add transports (e.g., POST method) if needed
	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.Use(gqlprometheus.NewTracer()) // Use as extension here

	// Conditionally enable introspection based on environment
	if os.Getenv("ENVIRONMENT") == "development" {
		graphqlHandler.Use(extension.Introspection{})
	}

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

	// Database health check endpoint
	app.Get("/db-health", func(c *fiber.Ctx) error {
		// Use the generated sqlc query method to check the database health
		err := queries.CheckHealth(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Database health check failed")
			return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
		}
		return c.SendString("Database is healthy")
	})

	// Prometheus route
	RegisterPrometheusRoute(app)
}
