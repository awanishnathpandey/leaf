package middleware

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/awanishnathpandey/leaf/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// MyClaims struct includes JWT standard claims and custom claims
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

// Global variables
var (
	queries                   *generated.Queries // Holds the database queries object
	userCache                 = make(map[string]cacheEntry)
	updateQueue               = make(chan string, 100) // Buffered channel for updates
	stopWorkers               = make(chan struct{})    // Channel to signal workers to stop
	wg                        sync.WaitGroup           // WaitGroup for worker synchronization
	cacheLock                 sync.RWMutex             // Protects userCache from concurrent access
	cacheTTL                  = 5 * time.Minute        // Cache expiration time
	cacheCleanupInterval      = 10 * time.Minute
	cacheMaxSize              = 1000
	numWorkers                = 10 // Number of worker goroutines
	unauthenticatedOperations = [][]byte{
		[]byte("login"),
		[]byte("register"),
		[]byte("refreshToken"),
		[]byte("forgotPassword"),
		[]byte("resetPassword"),
	}
)

// Cache entry structure
type cacheEntry struct {
	userID    int64
	userEmail string
	timestamp time.Time
}

// Initialize the queries object once
func InitializeQueries(q *generated.Queries) {
	queries = q
}

// StartWorkerPool initializes the worker pool
func StartWorkerPool() {
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker()
	}
	log.Info().Msgf("%d workers started", numWorkers)
	// Periodic cache cleanup
	go startCacheCleanup()
}

// StopWorkerPool gracefully stops all workers
func StopWorkerPool() {
	close(stopWorkers)
	wg.Wait()
	log.Info().Msg("All workers stopped")
}

// Worker goroutine processes tasks in the queue
func worker() {
	defer wg.Done()
	for {
		select {
		case userEmail := <-updateQueue:
			err := updateLastSeen(userEmail)
			if err != nil {
				log.Error().Err(err).Str("userEmail", userEmail).Msg("Failed to update last_seen")
			}
		case <-stopWorkers:
			return
		}
	}
}

// Update the database asynchronously
func updateLastSeen(userEmail string) error {
	ctx := context.Background()
	err := queries.UpdateUserLastSeenAtByEmail(ctx, userEmail)
	if err != nil {
		return fmt.Errorf("failed to update last_seen for user %s: %w", userEmail, err)
	}
	return nil
}

func startCacheCleanup() {
	ticker := time.NewTicker(cacheCleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cleanupExpiredCacheEntries()
			evictIfNeeded()
		case <-stopWorkers:
			return
		}
	}
}

func cleanupExpiredCacheEntries() {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	for email, entry := range userCache {
		if time.Since(entry.timestamp) > cacheTTL {
			delete(userCache, email)
		}
	}

	log.Info().Msg("User Cache cleanup complete")
}

// evictIfNeeded checks the cache size and evicts the oldest entry if the cache exceeds the max size.
func evictIfNeeded() {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	// If the cache exceeds the max size, evict the oldest entry
	if len(userCache) > cacheMaxSize {
		oldestEmail := ""
		oldestTimestamp := time.Now()

		// Find the oldest cache entry
		for email, entry := range userCache {
			if entry.timestamp.Before(oldestTimestamp) {
				oldestTimestamp = entry.timestamp
				oldestEmail = email
			}
		}

		// Delete the oldest entry from the cache
		if oldestEmail != "" {
			delete(userCache, oldestEmail)
			log.Info().Msgf("Evicted oldest cache entry: %s", oldestEmail)
		}
	}
}

// JWTMiddleware is a custom middleware to authenticate requests using JWT
func JWTMiddleware(queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the JWT_SECRET from the environment
		secretKey := []byte(os.Getenv("JWT_SECRET"))
		body := c.Body()

		// Check if the body contains any of the allowed operations
		for _, operation := range unauthenticatedOperations {
			if bytes.Contains(body, operation) {
				return c.Next()
			}
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

		// Check cache
		cacheLock.RLock()
		entry, found := userCache[userEmail]
		cacheLock.RUnlock()

		if found && time.Since(entry.timestamp) < cacheTTL {
			c.Locals("userID", entry.userID)
			c.Locals("userEmail", entry.userEmail)
			updateQueue <- entry.userEmail
			return c.Next()
		}

		// Query database if not in cache or expired
		ctx := context.Background()
		user, err := queries.GetUserByEmail(ctx, userEmail)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		cacheLock.Lock()
		userCache[userEmail] = cacheEntry{
			userID:    user.ID,
			userEmail: user.Email,
			timestamp: time.Now(),
		}
		cacheLock.Unlock()

		// Use Locals to store userID for easy access later in your handlers
		c.Locals("userID", user.ID)
		c.Locals("userEmail", user.Email)

		// Update last_seen in the database asynchronously by enqueueing the user ID
		updateQueue <- userEmail

		// Allow the request to continue to the next handler
		return c.Next()
	}
}
