package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
type MyRefreshClaims struct {
	ClientID string `json:"client_id"`
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Aud      string `json:"aud"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for a given user ID.
func GenerateJWT(userID int64, email, givenName, surname string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set in the environment")
	}

	clientID := os.Getenv("JWT_CLIENT_ID")
	if clientID == "" {
		return "", fmt.Errorf("JWT_CLIENT_ID is not set in the environment")
	}

	jwtIssuer := os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		return "", fmt.Errorf("JWT_ISSUER is not set in the environment")
	}

	jwtExpiryMinutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_MINUTES"))
	if err != nil {
		return "", fmt.Errorf("JWT_EXPIRY_MINUTES is invalid: %v", err)
	}

	jwtAudience := os.Getenv("JWT_AUDIENCE")
	if jwtAudience == "" {
		return "", fmt.Errorf("JWT_AUDIENCE is not set in the environment")
	}

	claims := &MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtExpiryMinutes) * time.Minute)),
			// IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer: jwtIssuer,
		},
		ClientID:  clientID,
		UID:       fmt.Sprintf("%d", userID),
		Mail:      email,
		GivenName: givenName,
		Surname:   surname,
		Email:     email,
		Aud:       jwtAudience,
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

// verifyWithOAuthToken sends the token to the verification endpoint
func VerifyWithOAuthToken(token string) (bool, error) {
	verificationURL := os.Getenv("JWT_VERIFICATION_URL")
	payload := map[string]string{
		"token":           token,
		"client_id":       os.Getenv("JWT_CLIENT_ID"),
		"token_type_hint": "access_token",
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return false, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, verificationURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("verification request failed: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to decode verification response: %w", err)
	}

	if active, ok := result["active"].(bool); ok {
		return active, nil
	}
	return false, fmt.Errorf("unexpected response format")
}

// GenerateJWT generates a JWT token for a given user ID.
func GenerateJWTRefresh(userID int64, email string) (string, error) {
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	if jwtRefreshSecret == "" {
		return "", fmt.Errorf("JWT_REFRESH_SECRET is not set in the environment")
	}

	clientID := os.Getenv("JWT_CLIENT_ID")
	if clientID == "" {
		return "", fmt.Errorf("JWT_CLIENT_ID is not set in the environment")
	}

	jwtIssuer := os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		return "", fmt.Errorf("JWT_ISSUER is not set in the environment")
	}

	jwtRefreshExpiryMinutes, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRY_MINUTES"))
	if err != nil {
		return "", fmt.Errorf("JWT_REFRESH_EXPIRY_MINUTES is invalid: %v", err)
	}

	jwtAudience := os.Getenv("JWT_AUDIENCE")
	if jwtAudience == "" {
		return "", fmt.Errorf("JWT_AUDIENCE is not set in the environment")
	}

	claims := &MyRefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtRefreshExpiryMinutes) * time.Minute)),
			// IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer: jwtIssuer,
		},
		ClientID: clientID,
		UID:      fmt.Sprintf("%d", userID),
		Email:    email,
		Aud:      jwtAudience,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(jwtRefreshSecret))
	if err != nil {
		return "", err
	}

	return signedRefreshToken, nil
}

// VerifyJWT verifies the provided JWT token and returns the claims
func VerifyJWTRefresh(tokenString string) (*jwt.RegisteredClaims, error) {
	refreshSecretKey := []byte(os.Getenv("JWT_REFRESH_SECRET"))

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key for verification
		return refreshSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// Return the claims (you can extract any custom claim here)
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("invalid refresh token claims")
	}
	return claims, nil
}
