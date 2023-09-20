package graph

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/wtkeqrf0/restService/api"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	*api.Server
}

func Validator(ctx context.Context, next graphql.Resolver) (any, error) {
	if v, ok := graphql.GetFieldContext(ctx).Args["req"].(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	return next(ctx)
}

func ErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	err = err.(*gqlerror.Error).Err

	return &gqlerror.Error{
		Err:     err,
		Message: "Failed to process request",
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code":  err.Error(),
			"error": err,
		},
		Rule: "",
	}
}
