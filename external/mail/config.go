package mail

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

type Config struct {
	SMTPHost       string
	SMTPPort       int
	SMTPUser       string
	SMTPPass       string
	SMTPFrom       string
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
}

// LoadConfig loads the SMTP configuration from environment variables.
func LoadConfig() (Config, error) {
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	from := os.Getenv("SMTP_FROM")
	connectTimeoutStr := os.Getenv("SMTP_CONNECT_TIMEOUT")
	sendTimeoutStr := os.Getenv("SMTP_SEND_TIMEOUT")

	// Validate required environment variables
	if host == "" || portStr == "" || user == "" || pass == "" || from == "" || connectTimeoutStr == "" || sendTimeoutStr == "" {
		return Config{}, fmt.Errorf("missing one or more required SMTP environment variables: SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASS, SMTP_CONNECT_TIMEOUT, SMTP_SEND_TIMEOUT")
	}

	// Convert SMTP_PORT to integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Error().Err(err).Msg("Invalid SMTP_PORT value")
		return Config{}, fmt.Errorf("invalid SMTP_PORT: %w", err)
	}

	// Convert SMTP_CONNECT_TIMEOUT to integer and then to time.Duration
	connectTimeout, err := strconv.Atoi(connectTimeoutStr)
	if err != nil {
		log.Error().Err(err).Msg("Invalid SMTP_CONNECT_TIMEOUT value")
		return Config{}, fmt.Errorf("invalid SMTP_CONNECT_TIMEOUT: %w", err)
	}

	// Convert SMTP_SEND_TIMEOUT to integer and then to time.Duration
	sendTimeout, err := strconv.Atoi(sendTimeoutStr)
	if err != nil {
		log.Error().Err(err).Msg("Invalid SMTP_SEND_TIMEOUT value")
		return Config{}, fmt.Errorf("invalid SMTP_SEND_TIMEOUT: %w", err)
	}

	return Config{
		SMTPHost:       host,
		SMTPPort:       port,
		SMTPUser:       user,
		SMTPPass:       pass,
		SMTPFrom:       from,
		ConnectTimeout: time.Duration(connectTimeout) * time.Second,
		SendTimeout:    time.Duration(sendTimeout) * time.Second,
	}, nil
}

// Example usage of the Config struct
func (cfg *Config) PrintConfig() {
	log.Info().
		Str("SMTPHost", cfg.SMTPHost).
		Int("SMTPPort", cfg.SMTPPort).
		Str("SMTPUser", cfg.SMTPUser).
		Str("SMTPFrom", cfg.SMTPFrom).
		Dur("ConnectTimeout", cfg.ConnectTimeout).
		Dur("SendTimeout", cfg.SendTimeout).
		Msg("SMTP Configuration Loaded")
}
