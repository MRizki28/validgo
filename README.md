# validgo

![Go Version](https://img.shields.io/badge/go-1.24%2B-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-success)

Laravel-style validation for Go Fiber.

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

func Register(c *fiber.Ctx) error {

    body, err := validgo.Parse[RegisterRequest](c)

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


    return c.JSON(fiber.Map{
        "status": true,
        "data":   body,
    })
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
* Gin support
* Echo support
* Query validation
* Form-data validation

---

# License

MIT

---

# Author

Muhammad Rizki
