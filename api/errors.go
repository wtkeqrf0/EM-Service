package api

import (
	"github.com/go-playground/validator/v10"
)

// ErrorType содержит типы возможных ошибок API.
type ErrorType string

const (
	ErrorValidation ErrorType = "validation"
	ErrorNotFound   ErrorType = "not_found"
	ErrorInternal   ErrorType = "internal"
)

// MyError структура описывает ошибку сервера.
type MyError struct {
	Type ErrorType `json:"type"`
}

func (a MyError) Error() string {
	switch a.Type {
	case ErrorValidation:
		return "validation_error"
	case ErrorInternal:
		return "internal_error"
	case ErrorNotFound:
		return "not_found_error"
	}

	return "unknown_error"
}

// ValidationError структура описывает ошибку валидации на уровне API.
type ValidationError struct {
	MyError `json:"-"`
	Fields  map[string]string `json:"fields"` // map[field]description.
}

// newValidationError создаёт ошибку валидации и
// заполняет её данными из полученных ошибок валидации.
func newValidationError(errs validator.ValidationErrors) *ValidationError {
	fields := make(map[string]string, len(errs))
	for _, err := range errs {
		fields[err.Field()] = err.Translate(eng)
	}

	return &ValidationError{
		MyError: MyError{Type: ErrorValidation},
		Fields:  fields,
	}
}

type AbstractError struct {
	MyError `json:"-"`
	Err     error `json:"error"`
}

func newError(err error, et ErrorType) AbstractError {
	return AbstractError{Err: err, MyError: MyError{Type: et}}
}
