package api

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/restService/pkg/ent"
)

type DeleteEnrichedFIORequest struct {
	ID int `json:"id" validate:"gt=0"`
}

type DeleteEnrichedFIOResponse struct {
	User *User `json:"user"`
}

func (r DeleteEnrichedFIORequest) Validate() error {
	err := vr.Struct(r)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	return newValidationError(errs)
}

func (s *Server) DeleteEnrichedFIO(ctx context.Context, r DeleteEnrichedFIORequest) (DeleteEnrichedFIOResponse, error) {
	user, err := s.ctrl.DeleteUser(ctx, r.ID)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return DeleteEnrichedFIOResponse{}, newError(err, ErrorNotFound)
		}
		return DeleteEnrichedFIOResponse{}, newError(err, ErrorInternal)
	}

	return DeleteEnrichedFIOResponse{User: &User{
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Age:        user.Age,
		Gender:     user.Gender,
		Country:    user.Country,
	}}, nil
}
