package api

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type TestFields struct {
	Number string `json:"phone" validate:"number"`
	MinAge *int   `json:"minAge,omitempty" validate:"omitempty,gte=0"`
	MaxAge *int   `json:"maxAge,omitempty" validate:"omitempty,gte=0,gtcsfield=MinAge"`
}

func TestValidator(t *testing.T) {
	v := newValidator(nil)

	age := 12

	err := v.Struct(&TestFields{MinAge: &age})
	errs := err.(validator.ValidationErrors)

	t.Log(errs.Translate(eng))
}
