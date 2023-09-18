package api

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/restService/internal/postgres"
	"github.com/wtkeqrf0/restService/pkg/ent"
)

type UpdateEnrichedFIORequest struct {
	EnrichedFio
}

type UpdateEnrichedFIOResponse struct {
	User *User
}

func (r UpdateEnrichedFIORequest) Validate() error {
	err := vr.Struct(r)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	return newValidationError(errs)
}

// UpdateEnrichedFIO updates FIO by ID from database.
// Returns an updated FIO.
func (s *Server) UpdateEnrichedFIO(ctx context.Context, r UpdateEnrichedFIORequest) (UpdateEnrichedFIOResponse, error) {
	user, err := s.ctrl.UpdateUser(ctx, postgres.UpdateEnrichedFIO(r.EnrichedFio))
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return UpdateEnrichedFIOResponse{}, newError(err, ErrorNotFound)
		}
		return UpdateEnrichedFIOResponse{}, newError(err, ErrorInternal)
	}

	return UpdateEnrichedFIOResponse{User: &User{
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Age:        user.Age,
		Gender:     user.Gender,
		Country:    user.Country,
	}}, err
}
