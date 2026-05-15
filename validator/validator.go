package validator

import (
	"strings"

	playground "github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *playground.Validate
}

func New() *Validator {
	v := playground.New()
	return &Validator{validate: v}
}

func (v *Validator) Validate(data any) map[string][]string {
	err := v.validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string][]string)
	for _, e := range err.(playground.ValidationErrors) {
		field := strings.ToLower(e.Field())
		errors[field] = append(errors[field], ParseMessage(e))
	}
	return errors
}
