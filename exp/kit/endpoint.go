package kit

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

// NewGeneric returns a new instance of Generic[req, res].
func NewGeneric[req, res any]() *Generic[req, res] {
	return &Generic[req, res]{}
}

// Generic is a generic endpoint implementation that can be used with any request and response types.
type Generic[req, res any] struct{}

// Endpoint returns an endpoint function that invokes the provided function e.
func (g Generic[req, res]) Endpoint(e func(context.Context, req) (res, error)) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return e(ctx, request.(req))
	}
}

// DecodeRequestFunc returns a function that decodes an HTTP request into a request object.
func (Generic[req, res]) DecodeRequestFunc(f func(context.Context, *http.Request) (req, error)) kithttp.DecodeRequestFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		return f(ctx, req)
	}
}

// EncodeResponseFunc returns a function that encodes a response object into an HTTP response.
func (Generic[req, res]) EncodeResponseFunc(f func(context.Context, http.ResponseWriter, res) error) kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, i interface{}) error {
		return f(ctx, w, i.(res))
	}
}
