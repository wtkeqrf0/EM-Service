package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/restService/pkg/ent"
)

// ErrorType contains types of possible API errors.
type ErrorType string

const (
	ErrorValidation ErrorType = "validation"
	ErrorNotFound   ErrorType = "not_found"
	ErrorInternal   ErrorType = "internal"
)

// MyError structure describes a server error.
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

// ValidationError the structure describes the validation error at the API level.
type ValidationError struct {
	MyError `json:"-"`
	Fields  map[string]string `json:"fields"` // map[field]description.
}

// newValidationError creates a validation error and
// fills it with data from the validation errors received.
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

// DBError the structure describes the error at the API level.
type DBError struct {
	MyError     `json:"-"`
	Description string `json:"description"`
}

// TODO: create a middleware to bring all errors to this type
func newDBError(err error) DBError {
	if _, ok := err.(*ent.NotFoundError); ok {
		return DBError{
			Description: err.Error()[5:],
			MyError: MyError{
				Type: ErrorNotFound,
			},
		}
	}

	return DBError{
		MyError: MyError{
			Type: ErrorInternal,
		},
		Description: err.Error(),
	}
}
