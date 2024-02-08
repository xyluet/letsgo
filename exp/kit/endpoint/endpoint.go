package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// GenericEndpoint is a higher-order function that wraps a given endpoint function and returns a new endpoint function.
type GenericEndpoint[Request any, Response any] func(context.Context, Request) (Response, error)

// Endpoint is a helper function to convert a GenericEndpoint to an endpoint.Endpoint.
// It is a helper function to convert a GenericEndpoint to an endpoint.Endpoint.
func Endpoint[Request, Response any](generic GenericEndpoint[Request, Response]) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return generic(ctx, request.(Request))
	}
}
