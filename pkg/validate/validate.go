package validate

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidator wraps the validator.Validate instance
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate validates the input struct
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

// New creates and returns a new CustomValidator instance
func New() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}
