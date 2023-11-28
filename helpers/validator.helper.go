package helpers

import "github.com/go-playground/validator"

type CustomValidator struct {
	Validator *validator.Validate
}

func (c CustomValidator) Validate(data interface{}) error {
	return c.Validator.Struct(data)
}
