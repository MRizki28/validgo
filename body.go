package validgo

import (
	"github.com/MRizki28/validgo/validator"
	"github.com/gofiber/fiber/v2"
)

func Parse[T any](c *fiber.Ctx) (T, error) {
	var body T

	if err := c.BodyParser(&body); err != nil {
		c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "invalid request body",
		})
		return body, err
	}

	v := validator.New()
	errors := v.Validate(body)

	if errors != nil {
		return body, c.Status(422).JSON(fiber.Map{
			"status":  false,
			"message": "validation failed",
			"errors":  errors,
		})
	}

	return body, nil
}
