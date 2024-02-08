package http_test

import (
	"context"
	"encoding/json"
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/stretchr/testify/assert"

	"github.com/xyluet/letsgo/exp/kit/transport/http"
)

func TestX(t *testing.T) {
	type requestType struct {
		Name string `json:"name"`
	}
	type responseType struct {
		Who string `json:"who"`
	}

	w, r := httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)
	http.NewServer[requestType, responseType](
		func(ctx context.Context, rt requestType) (responseType, error) {
			return responseType{Who: rt.Name}, nil
		},
		func(ctx context.Context, r *stdhttp.Request) (requestType, error) {
			return requestType{Name: "name"}, nil
		},
		func(ctx context.Context, w stdhttp.ResponseWriter, resp responseType) error {
			return json.NewEncoder(w).Encode(resp)
		},
		[]kithttp.ServerOption{}...,
	).ServeHTTP(w, r)
	assert.Equal(t, stdhttp.StatusOK, w.Code)
	assert.JSONEq(t, `{"who": "name"}`, w.Body.String())
}
