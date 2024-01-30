package http

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
)

// JSONResponseEncoder encodes the given response as JSON and writes it to the http.ResponseWriter.
// It uses the go-kit/kit/transport/http package to perform the encoding.
// If an error occurs during encoding, it returns an error wrapped with additional context information.
func JSONResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := kithttp.EncodeJSONResponse(ctx, w, response); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// WithStatusCode wraps the given response with a status code.
// It returns a new response object that implements the StatusCode() method to return the specified code.
func WithStatusCode(response interface{}, code int) interface{} {
	return statusCodeResponseWrapper{
		code:     code,
		response: response,
	}
}

type statusCodeResponseWrapper struct {
	code     int
	response interface{}
}

func (s statusCodeResponseWrapper) StatusCode() int {
	return s.code
}

func (s statusCodeResponseWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.response)
}
