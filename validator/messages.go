package validator

import (
	"fmt"
	"strings"

	playground "github.com/go-playground/validator/v10"
)

func ParseMessage(
	e playground.FieldError,
) string {
	field := strings.ToLower(e.Field())
	switch e.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", field, e.Param())
	default:
		return field + " is invalid"
	}
}