package api

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/wtkeqrf0/restService/internal/postgres"
)

type GetEnrichedFioRequest struct {
	Filter
}

type GetEnrichedFioResponse struct {
	Users []*User `json:"users"`
}

func (r GetEnrichedFioRequest) Validate() error {
	err := vr.Struct(r)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	return newValidationError(errs)
}

// GetEnrichedFio tries to get FIOs from database.
//
// If this function was caused by the same Request,
// return value will be retrieved from the cache.
func (s *Server) GetEnrichedFio(ctx context.Context, r GetEnrichedFioRequest) (GetEnrichedFioResponse, error) {

	key, err := s.ctrl.CacheKey(r)
	if err != nil {
		return GetEnrichedFioResponse{}, newError(err, ErrorInternal)
	}

	var resp GetEnrichedFioResponse
	if err = s.ctrl.Get(ctx, key, &resp); err != nil {
		return GetEnrichedFioResponse{}, newError(err, ErrorInternal)
	} else if resp.Users != nil {
		return resp, nil
	}

	users, err := s.ctrl.Users(ctx, postgres.Filter(r.Filter))
	if err != nil {
		return GetEnrichedFioResponse{}, newError(err, ErrorInternal)
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
		return GetEnrichedFioResponse{}, newError(err, ErrorInternal)
	}

	return resp, nil
}
