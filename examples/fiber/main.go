package main

import (
	"github.com/MRizki28/validgo"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func main() {
	app := fiber.New()

	app.Post(
		"/register",
		func(c *fiber.Ctx) error {
			var body RegisterRequest

			if err := c.BodyParser(&body); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":  false,
					"message": err.Error(),
				})
			}

			if err := validgo.Validate(body); err != nil {
				if valErr, ok := err.(*validgo.ValidationError); ok {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
						"status":  false,
						"message": "validation failed",
						"errors":  valErr.Errors,
					})
				}

				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":  false,
					"message": err.Error(),
				})
			}

			return c.Status(200).JSON(fiber.Map{
				"message": "success",
			})
		},
	)

	app.Listen(":3000")
}
