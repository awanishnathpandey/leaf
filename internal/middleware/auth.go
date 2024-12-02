package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// MyClaims struct that includes JWT standard claims and any custom claims
type MyClaims struct {
	ClientID  string `json:"client_id"`
	UID       string `json:"uid"`
	Mail      string `json:"mail"`
	GivenName string `json:"givenName"`
	Surname   string `json:"sn"`
	Email     string `json:"email"`
	Aud       string `json:"aud"`
	jwt.RegisteredClaims
}

// Global variable to hold queries object
var queries *generated.Queries

// Initialize the queries object once at the beginning of the app
func InitializeQueries(q *generated.Queries) {
	queries = q
}

// Cache for user existence checks (key = userEmail, value = cacheEntry)
var userCache = make(map[string]cacheEntry)

// Cache entry structure, including the timestamp of when it was created
type cacheEntry struct {
	userID    int64
	userEmail string
	timestamp time.Time
}

// Cache expiration time (set to 5 minutes)
const cacheExpiration = 5 * time.Minute

// Asynchronous worker to update last_seen
const numWorkers = 10                    // Number of workers in the pool
var updateQueue = make(chan string, 100) // Buffered channel for user IDs

func init() {
	// Start worker goroutines for processing updates
	for i := 0; i < numWorkers; i++ {
		go worker()
	}
}

// Worker to process updates
func worker() {
	for userEmail := range updateQueue {
		err := updateLastSeen(userEmail)
		if err != nil {
			log.Error().Err(err).Str("userEmail", userEmail).Msg("Failed to update last_seen")
		}
	}
}

// Update the database asynchronously
func updateLastSeen(userEmail string) error {
	ctx := context.Background()

	// Update the last_seen in the database for the given userEmail
	err := queries.UpdateUserLastSeenAtByEmail(ctx, userEmail)
	if err != nil {
		return fmt.Errorf("failed to update last_seen for user %s: %w", userEmail, err)
	}
	return nil
}

// JWTMiddleware is a custom middleware to authenticate requests using JWT
func JWTMiddleware(queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the JWT_SECRET from the environment
		secretKey := []byte(os.Getenv("JWT_SECRET"))

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
		_, _, err := jwt.NewParser().ParseUnverified(tokenString, claims)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Invalid token",
				"details": err.Error(),
			})
		}

		// log.Info().
		// 	Str("Issuer", claims.Issuer).
		// 	Str("ClientID", claims.ClientID).
		// 	Str("UID", claims.UID).
		// 	Msg("Parsed token claims")
		switch claims.Issuer {
		case "leaf":
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
		case "different":
			// Validate token with the OAuth verification endpoint
			isValid, err := utils.VerifyWithOAuthToken(tokenString)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   "OAuth Token verification failed for OAuth",
					"details": err.Error(),
				})
			}
			// Check if the OAuth token is valid
			if !isValid {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid OAuthtoken",
				})
			}

		default:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unknown issuer",
			})
		}

		// Check if the verified token is expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token expired",
			})
		}

		// Store claims (user info, etc.) in the request context
		userEmail := claims.Email // Retrieve userEmail from claims
		if userEmail == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Email not found in token",
			})
		}

		// Check cache for user existence with expiration handling
		if entry, found := userCache[userEmail]; found {
			if time.Since(entry.timestamp) < cacheExpiration {
				// Valid cache entry
				c.Locals("userID", entry.userID)
				c.Locals("userEmail", entry.userEmail)
				updateQueue <- entry.userEmail // Enqueue userEmail for last_seen update
				return c.Next()
			}
			// Cache expired, remove entry
			delete(userCache, userEmail)
		}

		// Query the database if the user is not in the cache or cache expired
		ctx := context.Background()
		user, err := queries.GetUserByEmail(ctx, userEmail)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		// Update cache
		userCache[userEmail] = cacheEntry{
			userID:    user.ID,
			userEmail: user.Email,
			timestamp: time.Now(),
		}

		// Use Locals to store userID for easy access later in your handlers
		c.Locals("userID", user.ID)
		c.Locals("userEmail", user.Email)

		// Update last_seen in the database asynchronously by enqueueing the user ID
		updateQueue <- userEmail

		// Allow the request to continue to the next handler
		return c.Next()
	}
}
