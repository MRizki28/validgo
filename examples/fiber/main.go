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
			body, err := validgo.Parse[RegisterRequest](c)
			if err != nil {
				if c.BodyParser(&body) != nil {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"status":  false,
						"message": "invalid request body",
					})
				}

				if valErr, ok := err.(*validgo.ValidationError); ok {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
						"status":  false,
						"message": "validation failed",
						"errors":  valErr.Errors,
					})
				}

				return err
			}

			return c.JSON(fiber.Map{
				"status": true,
				"data":   body,
			})
		},
	)

	app.Listen(":3000")
}
