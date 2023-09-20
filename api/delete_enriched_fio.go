package api

import (
	"context"
	"github.com/go-playground/validator/v10"
)

type DeleteEnrichedFioRequest struct {
	ID int `json:"id" validate:"gt=0"`
}

type DeleteEnrichedFioResponse struct {
	User *User `json:"user,omitempty"`
}

func (r DeleteEnrichedFioRequest) Validate() error {
	err := vr.Struct(r)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	return newValidationError(errs)
}

// DeleteEnrichedFio deletes the FIO from database by id.
// Returns deleted FIO.
func (s *Server) DeleteEnrichedFio(ctx context.Context, r DeleteEnrichedFioRequest) (DeleteEnrichedFioResponse, error) {
	user, err := s.ctrl.DeleteUser(ctx, r.ID)
	if err != nil {
		return DeleteEnrichedFioResponse{}, newDBError(err)
	}

	return DeleteEnrichedFioResponse{User: &User{
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Age:        user.Age,
		Gender:     user.Gender,
		Country:    user.Country,
	}}, nil
}
