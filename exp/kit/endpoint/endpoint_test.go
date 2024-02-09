package endpoint_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xyluet/letsgo/exp/kit/endpoint"
)

func TestGeneric(t *testing.T) {
	type requestType struct {
		Name string `json:"name"`
	}

	type responseType struct {
		Name string `json:"name"`
		What string `json:"what"`
	}

	anEndpoint := func(ctx context.Context, req requestType) (responseType, error) {
		return responseType{Name: req.Name, What: "foo"}, nil
	}
	endpoint := endpoint.Generic[requestType, responseType](anEndpoint).Endpoint()
	resp, err := endpoint(context.Background(), requestType{Name: "a name"})
	assert.NoError(t, err)
	assert.IsType(t, responseType{}, resp)
	assert.Panics(t, func() {
		endpoint(context.Background(), "not a request type")
	})
}
