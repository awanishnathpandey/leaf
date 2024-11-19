package middleware

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// MyClaims struct that includes JWT standard claims and any custom claims
type MyClaims struct {
	jwt.RegisteredClaims
	CustomField string `json:"custom_field"`
}

// Global variable to hold queries object
var queries *generated.Queries

// Initialize the queries object once at the beginning of the app
func InitializeQueries(q *generated.Queries) {
	queries = q
}

// Cache for user existence checks (key = userID, value = cacheEntry)
var userCache = make(map[int64]cacheEntry)

// Cache entry structure, including the timestamp of when it was created
type cacheEntry struct {
	exists    bool
	timestamp time.Time
}

// Cache expiration time (set to 5 minutes)
const cacheExpiration = 5 * time.Minute

// Asynchronous worker to update last_seen
const numWorkers = 10                   // Number of workers in the pool
var updateQueue = make(chan int64, 100) // Buffered channel for user IDs

func init() {
	// Start worker goroutines for processing updates
	for i := 0; i < numWorkers; i++ {
		go worker()
	}
}

// Worker to process updates
func worker() {
	for userID := range updateQueue {
		err := updateLastSeen(userID)
		if err != nil {
			log.Error().Err(err).Int64("userID", userID).Msg("Failed to update last_seen")
			// log.Printf("Failed to update last_seen for user %d: %v\n", userID, err)
		}
	}
}

// Update the database asynchronously
func updateLastSeen(userID int64) error {
	ctx := context.Background()

	// Update the last_seen in the database
	err := queries.UpdateUserLastSeenAt(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to update last_seen for user %d: %w", userID, err)
	}
	return nil
}

// JWTMiddleware is a custom middleware to authenticate requests using JWT
func JWTMiddleware(queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the JWT_SECRET from the environment
		secretKey := []byte(os.Getenv("JWT_SECRET"))

		// Skip authentication for login/register routes
		// if strings.Contains(c.Path(), "login") || strings.Contains(c.Path(), "register") {
		// 	return c.Next()
		// }
		body := c.Body()
		if strings.Contains(string(body), "login") || strings.Contains(string(body), "register") {
			return c.Next()
		}

		// Get the JWT token from the request's Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		// Validate that the token starts with "Bearer " prefix
		if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}
		tokenString := authHeader[7:] // Extract token after "Bearer "

		// Ensure the token is well-formed (has 3 parts)
		parts := strings.Split(tokenString, ".")
		if len(parts) != 3 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}

		// Parse and validate the token
		claims := &MyClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Check if the signing method matches
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Invalid token",
				"details": err.Error(),
			})
		}

		// Check if the token is valid
		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Check if the token is expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token expired",
			})
		}

		// Store claims (user info, etc.) in the request context
		userID := claims.Subject // The subject is typically a string
		uidInt64, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid user ID in token",
			})
		}

		// Check cache for user existence with expiration handling
		if entry, found := userCache[uidInt64]; found {
			// If the cache entry exists and is not expired, proceed
			if time.Since(entry.timestamp) < cacheExpiration {
				if !entry.exists {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"error": "User not found",
					})
				}
			} else {
				// Cache expired, remove the entry and recheck
				delete(userCache, uidInt64)
			}
		}

		// Query the database if the user is not in the cache or cache expired
		ctx := context.Background()
		_, err = queries.GetUser(ctx, uidInt64)
		if err != nil {
			// Cache the result as user not found
			userCache[uidInt64] = cacheEntry{
				exists:    false,
				timestamp: time.Now(),
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		// User found, update cache
		userCache[uidInt64] = cacheEntry{
			exists:    true,
			timestamp: time.Now(),
		}

		// Update last_seen in the database asynchronously by enqueueing the user ID
		updateQueue <- uidInt64

		// Use Locals to store userID for easy access later in your handlers
		c.Locals("userID", uidInt64)

		// Allow the request to continue to the next handler
		return c.Next()
	}
}
