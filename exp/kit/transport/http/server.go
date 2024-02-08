package http

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/xyluet/letsgo/exp/kit/endpoint"
)

// NewServer returns a new kithttp.Server.
// It is a helper function to convert a GenericEndpoint to a kithttp.Server.
func NewServer[Request, Response any](
	e endpoint.GenericEndpoint[Request, Response],
	dec func(context.Context, *http.Request) (Request, error),
	enc func(context.Context, http.ResponseWriter, Response) error,
	options ...kithttp.ServerOption,
) *kithttp.Server {
	return kithttp.NewServer(
		endpoint.Endpoint(e),
		func(ctx context.Context, r *http.Request) (request interface{}, err error) {
			return dec(ctx, r)
		},
		func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			return enc(ctx, w, response.(Response))
		},
		options...,
	)
}
