package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/MRizki28/validgo"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func TestValidationFailed(t *testing.T) {
	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
	var body RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	err := validgo.Validate(body)

	if err != nil {
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
})

	body := map[string]any{
		"name":  "",
		"email": "invalid-email",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(
		"POST",
		"/register",
		bytes.NewReader(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	responseBody, _ := io.ReadAll(resp.Body)

	// Parse response untuk validasi lebih detail
	var response map[string]any
	json.Unmarshal(responseBody, &response)

	t.Logf("Response: %+v", response)

	if resp.StatusCode != 422 {
		t.Errorf("expected status 422 got %d", resp.StatusCode)
	}

	if response["status"] != false {
		t.Error("expected status false")
	}

	if response["errors"] == nil {
		t.Error("expected errors field to be present")
	}
}