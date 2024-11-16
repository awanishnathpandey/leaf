package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a JWT token for a given user ID.
func GenerateJWT(userID int64) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set in the environment")
	}

	jwtIssuer := os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		return "", fmt.Errorf("JWT_ISSUER is not set in the environment")
	}

	jwtExpiryMinutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_MINUTES"))
	if err != nil {
		return "", fmt.Errorf("JWT_EXPIRY_MINUTES is invalid: %v", err)
	}

	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtExpiryMinutes) * time.Minute)), // 1-day expiry
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    jwtIssuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyJWT verifies the provided JWT token and returns the claims
func VerifyJWT(tokenString string) (*jwt.RegisteredClaims, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key for verification
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Return the claims (you can extract any custom claim here)
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
