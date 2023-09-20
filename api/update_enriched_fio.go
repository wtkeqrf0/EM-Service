package api

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/restService/internal/postgres"
	"github.com/wtkeqrf0/restService/pkg/ent"
)

type UpdateEnrichedFioRequest struct {
	EnrichedFio
}

type UpdateEnrichedFioResponse struct {
	User *User `json:"user,omitempty"`
}

func (r UpdateEnrichedFioRequest) Validate() error {
	err := vr.Struct(r)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	return newValidationError(errs)
}

// UpdateEnrichedFio updates FIO by ID from database.
// Returns an updated FIO.
func (s *Server) UpdateEnrichedFio(ctx context.Context, r UpdateEnrichedFioRequest) (UpdateEnrichedFioResponse, error) {
	user, err := s.ctrl.UpdateUser(ctx, postgres.UpdateEnrichedFIO(r.EnrichedFio))
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return UpdateEnrichedFioResponse{}, newError(err, ErrorNotFound)
		}
		return UpdateEnrichedFioResponse{}, newError(err, ErrorInternal)
	}

	return UpdateEnrichedFioResponse{User: &User{
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Age:        user.Age,
		Gender:     user.Gender,
		Country:    user.Country,
	}}, err
}
