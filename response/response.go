package response

import "github.com/gofiber/fiber/v2"

func ValidationError(
	errors any,
) fiber.Map {
	return fiber.Map{
		"status": false,
		"message": "validation failed",
		"errors": errors,
	}
}