package http

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/xyluet/letsgo/exp/kit/endpoint"
)

// NewServer returns a new http.Handler that wraps the provided endpoint.
func NewServer[Request, Response any](
	g endpoint.Generic[Request, Response],
	dec func(context.Context, *http.Request) (Request, error),
	enc func(context.Context, http.ResponseWriter, Response) error,
	options ...kithttp.ServerOption,
) *kithttp.Server {
	return kithttp.NewServer(
		g.Endpoint(),
		func(ctx context.Context, r *http.Request) (request interface{}, err error) {
			return dec(ctx, r)
		},
		func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			return enc(ctx, w, response.(Response))
		},
		options...,
	)
}
