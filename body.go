package validgo

import (
	"github.com/MRizki28/validgo/validator"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Errors interface{}
}

func (e *ValidationError) Error() string {
	return "validation failed"
}

func Parse[T any](c *fiber.Ctx) (T, error) {
	var body T

	if err := c.BodyParser(&body); err != nil {
		return body, err
	}

	v := validator.New()
	errors := v.Validate(body)

	if errors != nil {
		return body, &ValidationError{Errors: errors}
	}

	return body, nil
}