package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

// FormatValidationErrors takes a validation error and returns a user-friendly message
func FormatValidationErrors(err error) error {
	var validationErrors []string
	for _, e := range err.(validator.ValidationErrors) {
		// Customizing error messages based on the field and tag
		switch e.Tag() {
		case "required":
			return fmt.Errorf("%s is required", e.Field())
		case "min":
			return fmt.Errorf("%s must be at least %s characters", e.Field(), e.Param())
		case "max":
			return fmt.Errorf("%s must not exceed %s characters", e.Field(), e.Param())
		case "len":
			return fmt.Errorf("%s must be exactly %s characters", e.Field(), e.Param())
		case "eq":
			return fmt.Errorf("%s must be equal to %s", e.Field(), e.Param())
		case "ne":
			return fmt.Errorf("%s must not be equal to %s", e.Field(), e.Param())
		case "lt":
			return fmt.Errorf("%s must be less than %s", e.Field(), e.Param())
		case "lte":
			return fmt.Errorf("%s must be less than or equal to %s", e.Field(), e.Param())
		case "gt":
			return fmt.Errorf("%s must be greater than %s", e.Field(), e.Param())
		case "gte":
			return fmt.Errorf("%s must be greater than or equal to %s", e.Field(), e.Param())
		case "alpha":
			return fmt.Errorf("%s must contain only alphabetic characters", e.Field())
		case "alphanum":
			return fmt.Errorf("%s must contain only alphanumeric characters", e.Field())
		case "numeric":
			return fmt.Errorf("%s must be a numeric value", e.Field())
		case "boolean":
			return fmt.Errorf("%s must be a boolean value", e.Field())
		case "email":
			return fmt.Errorf("%s must be a valid email address", e.Field())
		case "url":
			return fmt.Errorf("%s must be a valid URL", e.Field())
		case "uuid":
			return fmt.Errorf("%s must be a valid UUID", e.Field())
		case "contains":
			return fmt.Errorf("%s must contain the substring '%s'", e.Field(), e.Param())
		case "startswith":
			return fmt.Errorf("%s must start with '%s'", e.Field(), e.Param())
		case "endswith":
			return fmt.Errorf("%s must end with '%s'", e.Field(), e.Param())
		case "excluded_with":
			return fmt.Errorf("%s cannot be present when %s is present", e.Field(), e.Param())
		case "excluded_without":
			return fmt.Errorf("%s cannot be present without %s", e.Field(), e.Param())
		case "oneof":
			return fmt.Errorf("%s must be one of [%s]", e.Field(), e.Param())
		case "containsany":
			return fmt.Errorf("%s must contain at least one of [%s]", e.Field(), e.Param())
		case "excludes":
			return fmt.Errorf("%s must not contain '%s'", e.Field(), e.Param())
		case "excludesall":
			return fmt.Errorf("%s must not contain any of [%s]", e.Field(), e.Param())
		case "excludesrune":
			return fmt.Errorf("%s must not contain the rune '%s'", e.Field(), e.Param())
		case "iso3166_1_alpha2":
			return fmt.Errorf("%s must be a valid ISO3166-1 alpha-2 country code", e.Field())
		case "iso3166_1_alpha3":
			return fmt.Errorf("%s must be a valid ISO3166-1 alpha-3 country code", e.Field())
		case "iso3166_1_numeric":
			return fmt.Errorf("%s must be a valid ISO3166-1 numeric country code", e.Field())
		case "iso4217":
			return fmt.Errorf("%s must be a valid ISO4217 currency code", e.Field())
		case "ip":
			return fmt.Errorf("%s must be a valid IP address", e.Field())
		case "ipv4":
			return fmt.Errorf("%s must be a valid IPv4 address", e.Field())
		case "ipv6":
			return fmt.Errorf("%s must be a valid IPv6 address", e.Field())
		case "mac":
			return fmt.Errorf("%s must be a valid MAC address", e.Field())
		case "hostname":
			return fmt.Errorf("%s must be a valid hostname", e.Field())
		case "fqdn":
			return fmt.Errorf("%s must be a valid fully qualified domain name", e.Field())
		case "unique":
			return fmt.Errorf("%s must contain unique values", e.Field())
		case "ascii":
			return fmt.Errorf("%s must contain only ASCII characters", e.Field())
		case "printascii":
			return fmt.Errorf("%s must contain only printable ASCII characters", e.Field())
		case "base64":
			return fmt.Errorf("%s must be a valid Base64 encoded string", e.Field())
		case "hexadecimal":
			return fmt.Errorf("%s must be a valid hexadecimal value", e.Field())
		case "rgb":
			return fmt.Errorf("%s must be a valid RGB color", e.Field())
		case "rgba":
			return fmt.Errorf("%s must be a valid RGBA color", e.Field())
		case "latitude":
			return fmt.Errorf("%s must be a valid latitude coordinate", e.Field())
		case "longitude":
			return fmt.Errorf("%s must be a valid longitude coordinate", e.Field())
		case "timezone":
			return fmt.Errorf("%s must be a valid timezone", e.Field())
		default:
			return fmt.Errorf("%s is invalid or does not meet the expected criteria", e.Field())
		}
	}
	return fmt.Errorf("validation failed: %v", validationErrors)
}
