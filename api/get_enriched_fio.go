package api

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/restService/internal/postgres"
)

type GetEnrichedFIORequest struct {
	Filter
}

type GetEnrichedFIOResponse struct {
	Users []*User `json:"users"`
}

func (r GetEnrichedFIORequest) Validate() error {
	err := vr.Struct(r)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	return newValidationError(errs)
}

func (s *Server) GetEnrichedFIO(ctx context.Context, r GetEnrichedFIORequest) (GetEnrichedFIOResponse, error) {

	key, err := s.ctrl.CacheKey(r)
	if err != nil {
		return GetEnrichedFIOResponse{}, newError(err, ErrorInternal)
	}

	var resp GetEnrichedFIOResponse
	if err = s.ctrl.Get(ctx, key, &resp); err != nil {
		return GetEnrichedFIOResponse{}, newError(err, ErrorInternal)
	} else if resp.Users != nil {
		return resp, nil
	}

	users, err := s.ctrl.Users(ctx, postgres.Filter(r.Filter))
	if err != nil {
		return GetEnrichedFIOResponse{}, newError(err, ErrorInternal)
	}

	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID:         user.ID,
			Name:       user.Name,
			Surname:    user.Surname,
			Patronymic: user.Patronymic,
			Age:        user.Age,
			Gender:     user.Gender,
			Country:    user.Country,
		}
	}

	resp.Users = res

	if err = s.ctrl.Save(ctx, key, resp); err != nil {
		return GetEnrichedFIOResponse{}, newError(err, ErrorInternal)
	}

	return resp, nil
}
