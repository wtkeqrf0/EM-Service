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
			graphql.AddError(ctx, &gqlerror.Error{
				Message: "Request is not valid",
				Path:    graphql.GetPath(ctx),
				Extensions: map[string]interface{}{
					"code":  err.Error(),
					"error": err,
				},
			})
			return nil, nil
		}
	}

	return next(ctx)
}
