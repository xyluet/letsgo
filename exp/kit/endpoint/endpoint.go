package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Generic is a helper function to convert a function to an endpoint.Endpoint.
type Generic[Request any, Response any] func(context.Context, Request) (Response, error)

// Endpoint returns an endpoint.Endpoint.
func (g Generic[Request, Response]) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return g(ctx, request.(Request))
	}
}
