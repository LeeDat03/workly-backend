package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator is a wrapper around go-playground/validator
type Validator struct {
	validate *validator.Validate
}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// Validate validates a struct
func (v *Validator) Validate(data interface{}) error {
	return v.validate.Struct(data)
}

// ValidateField validates a single field
func (v *Validator) ValidateField(field interface{}, tag string) error {
	return v.validate.Var(field, tag)
}

// GetValidate returns the underlying validator instance
func (v *Validator) GetValidate() *validator.Validate {
	return v.validate
}
