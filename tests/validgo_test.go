package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	validgo "github.com/MRizki28/validgo"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func TestValidationFailed(t *testing.T) {

	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {

		_, err := validgo.Parse[RegisterRequest](c)

		return err
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

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	responseBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(responseBody))

	if resp.StatusCode != 422 {
		t.Errorf(
			"expected status 422 got %d",
			resp.StatusCode,
		)
	}
}