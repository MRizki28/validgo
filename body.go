package validgo

import "github.com/MRizki28/validgo/validator"

type ValidationError struct {
	Errors interface{}
}

func(e *ValidationError) Error() string {
	return "Validation Error"
}

func Validate[T any](data T) error {
	v := validator.New()
	errors := v.Validate(data)

	if errors != nil {
		return &ValidationError{Errors: errors}
	}

	return nil
}