# validgo

![Go Version](https://img.shields.io/badge/go-1.24%2B-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-success)

Validation for Golang.

`validgo` helps simplify request parsing and validation in Go by wrapping `go-playground/validator` with a cleaner developer experience.

## Features

* Request body parsing
* Automatic validation
* Automatic JSON validation response
* Generic typed request parsing
* Clean Fiber integration
* Built on top of go-playground/validator

---

# Installation

```bash
go get github.com/MRizki28/validgo
```

Install validgo CLI:

```bash
go install github.com/MRizki28/validgo/cmd/validgo@latest
```

Make sure your Go binary path is added to PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
# CLI Commands

## Generate Request

```bash
validgo make:request RegisterUser
```

Generated file:

```bash
app/requests/register_user_request.go
```

---

# Quick Start

## Request Struct

```go
package main

type RegisterRequest struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}
```

---

## Example

```go
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

```

---


# Validation Error Response

```json
{
  "status": false,
  "message": "validation failed",
  "errors": {
    "email": [
      "email must be a valid email"
    ],
    "name": [
      "name is required"
    ]
  }
}
```

---

# Invalid Body Response

```json
{
  "status": false,
  "message": "invalid request body"
}
```

---

# Example

```bash
go run examples/fiber/main.go
```

---

# Testing

Run tests:

```bash
go test -v ./...
```

---

# Roadmap

* Custom validation messages
* Localization support
* Database unique validation
* Snake case field formatter
* Query validation
* Form-data validation

---

# License

MIT

---

# Author

Muhammad Rizki
