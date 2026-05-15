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

	app.Post("/register", func(c *fiber.Ctx) error {

		body, err := validgo.Parse[RegisterRequest](c)

		if err != nil {

			if ve, ok := err.(validgo.ValidationError); ok {
				return c.Status(422).JSON(fiber.Map{
					"status":  false,
					"message": "validation failed",
					"errors":  ve.Errors,
				})
			}

			return c.Status(400).JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"status": true,
			"data":   body,
		})
	})

	app.Listen(":3000")
}
