package validations

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func validateFields(fields map[string]struct {
	Value interface{}
	Tag   string
}) error {
	var messages []string

	for fieldName, field := range fields {
		err := validate.Var(field.Value, field.Tag)
		if err != nil {
			messages = append(messages, formatValidationError(fieldName, err))
		}
	}

	if len(messages) > 0 {
		return errors.New(strings.Join(messages, " "))
	}
	return nil
}

func formatValidationError(fieldName string, err error) string {
	var messages []string

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return fmt.Sprintf("Field '%s' has an invalid validation setup.", fieldName)
	}

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			messages = append(messages, fmt.Sprintf("%s is required.", fieldName))
		case "min":
			messages = append(messages, fmt.Sprintf("%s must be at least %s characters.", fieldName, e.Param()))
		case "max":
			messages = append(messages, fmt.Sprintf("%s must not exceed %s characters.", fieldName, e.Param()))
		case "len":
			messages = append(messages, fmt.Sprintf("%s must be exactly %s characters.", fieldName, e.Param()))
		case "eq":
			messages = append(messages, fmt.Sprintf("%s must be equal to %s.", fieldName, e.Param()))
		case "ne":
			messages = append(messages, fmt.Sprintf("%s must not be equal to %s.", fieldName, e.Param()))
		case "lt":
			messages = append(messages, fmt.Sprintf("%s must be less than %s.", fieldName, e.Param()))
		case "lte":
			messages = append(messages, fmt.Sprintf("%s must be less than or equal to %s.", fieldName, e.Param()))
		case "gt":
			messages = append(messages, fmt.Sprintf("%s must be greater than %s.", fieldName, e.Param()))
		case "gte":
			messages = append(messages, fmt.Sprintf("%s must be greater than or equal to %s.", fieldName, e.Param()))
		case "alpha":
			messages = append(messages, fmt.Sprintf("%s must contain only alphabetic characters.", fieldName))
		case "alphanum":
			messages = append(messages, fmt.Sprintf("%s must contain only alphanumeric characters.", fieldName))
		case "numeric":
			messages = append(messages, fmt.Sprintf("%s must be a numeric value.", fieldName))
		case "boolean":
			messages = append(messages, fmt.Sprintf("%s must be a boolean value.", fieldName))
		case "email":
			messages = append(messages, fmt.Sprintf("%s must be a valid email address.", fieldName))
		case "url":
			messages = append(messages, fmt.Sprintf("%s must be a valid URL.", fieldName))
		case "uuid":
			messages = append(messages, fmt.Sprintf("%s must be a valid UUID.", fieldName))
		case "contains":
			messages = append(messages, fmt.Sprintf("%s must contain the substring '%s'.", fieldName, e.Param()))
		case "startswith":
			messages = append(messages, fmt.Sprintf("%s must start with '%s'.", fieldName, e.Param()))
		case "endswith":
			messages = append(messages, fmt.Sprintf("%s must end with '%s'.", fieldName, e.Param()))
		case "excluded_with":
			messages = append(messages, fmt.Sprintf("%s cannot be present when %s is present.", fieldName, e.Param()))
		case "excluded_without":
			messages = append(messages, fmt.Sprintf("%s cannot be present without %s.", fieldName, e.Param()))
		case "oneof":
			messages = append(messages, fmt.Sprintf("%s must be one of [%s].", fieldName, e.Param()))
		case "containsany":
			messages = append(messages, fmt.Sprintf("%s must contain at least one of [%s].", fieldName, e.Param()))
		case "excludes":
			messages = append(messages, fmt.Sprintf("%s must not contain '%s'.", fieldName, e.Param()))
		case "excludesall":
			messages = append(messages, fmt.Sprintf("%s must not contain any of [%s].", fieldName, e.Param()))
		case "excludesrune":
			messages = append(messages, fmt.Sprintf("%s must not contain the rune '%s'.", fieldName, e.Param()))
		case "iso3166_1_alpha2":
			messages = append(messages, fmt.Sprintf("%s must be a valid ISO3166-1 alpha-2 country code.", fieldName))
		case "iso3166_1_alpha3":
			messages = append(messages, fmt.Sprintf("%s must be a valid ISO3166-1 alpha-3 country code.", fieldName))
		case "iso3166_1_numeric":
			messages = append(messages, fmt.Sprintf("%s must be a valid ISO3166-1 numeric country code.", fieldName))
		case "iso4217":
			messages = append(messages, fmt.Sprintf("%s must be a valid ISO4217 currency code.", fieldName))
		case "ip":
			messages = append(messages, fmt.Sprintf("%s must be a valid IP address.", fieldName))
		case "ipv4":
			messages = append(messages, fmt.Sprintf("%s must be a valid IPv4 address.", fieldName))
		case "ipv6":
			messages = append(messages, fmt.Sprintf("%s must be a valid IPv6 address.", fieldName))
		case "mac":
			messages = append(messages, fmt.Sprintf("%s must be a valid MAC address.", fieldName))
		case "hostname":
			messages = append(messages, fmt.Sprintf("%s must be a valid hostname.", fieldName))
		case "fqdn":
			messages = append(messages, fmt.Sprintf("%s must be a valid fully qualified domain name.", fieldName))
		case "unique":
			messages = append(messages, fmt.Sprintf("%s must contain unique values.", fieldName))
		case "ascii":
			messages = append(messages, fmt.Sprintf("%s must contain only ASCII characters.", fieldName))
		case "printascii":
			messages = append(messages, fmt.Sprintf("%s must contain only printable ASCII characters.", fieldName))
		case "base64":
			messages = append(messages, fmt.Sprintf("%s must be a valid Base64 encoded string.", fieldName))
		case "hexadecimal":
			messages = append(messages, fmt.Sprintf("%s must be a valid hexadecimal value.", fieldName))
		case "rgb":
			messages = append(messages, fmt.Sprintf("%s must be a valid RGB color.", fieldName))
		case "rgba":
			messages = append(messages, fmt.Sprintf("%s must be a valid RGBA color.", fieldName))
		case "latitude":
			messages = append(messages, fmt.Sprintf("%s must be a valid latitude coordinate.", fieldName))
		case "longitude":
			messages = append(messages, fmt.Sprintf("%s must be a valid longitude coordinate.", fieldName))
		case "timezone":
			messages = append(messages, fmt.Sprintf("%s must be a valid timezone.", fieldName))
		default:
			messages = append(messages, fmt.Sprintf("field '%s' failed validation: failed on the '%s' tag.", fieldName, e.Tag()))
		}
	}
	return strings.Join(messages, " ")
}
