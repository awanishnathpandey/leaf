package middleware

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/awanishnathpandey/leaf/db/generated"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// MyClaims struct that includes JWT standard claims and any custom claims
type MyClaims struct {
	jwt.RegisteredClaims
	CustomField string `json:"custom_field"`
}

// JWTMiddleware is a custom middleware to authenticate requests using JWT
func JWTMiddleware(queries *generated.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Print loaded JWT_SECRET to debug
		secretKey := []byte(os.Getenv("JWT_SECRET"))
		// fmt.Println("JWT_SECRET from environment:", secretKey)
		// fmt.Println("JWT Secret Key: ", string(secretKey))
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
		tokenString := authHeader
		if len(tokenString) <= 7 || tokenString[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}
		tokenString = tokenString[7:] // Extract token after "Bearer "

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

		// Check if user exists
		ctx := context.Background()
		_, err = queries.GetUser(ctx, uidInt64)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		// Update last_seen in the database
		if err := queries.UpdateUserLastSeenAt(ctx, uidInt64); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update last seen",
			})
		}

		// Use Locals to store userID for easy access later in your handlers
		c.Locals("userID", uidInt64)

		// Allow the request to continue to the next handler
		return c.Next()
	}
}
