package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/stretchr/testify/assert"
	"github.com/xyluet/letsgo/kitext/transport/http"
)

func TestJSONResponseEncoder(t *testing.T) {
	type resp struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}

	response := resp{ID: "id", Text: "text"}

	rec := httptest.NewRecorder()
	err := http.JSONResponseEncoder(context.Background(), rec, response)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"id":"id","text":"text"}`, rec.Body.String())
}

type marshalJSONError struct{}

func (m marshalJSONError) MarshalJSON() ([]byte, error) {
	return nil, assert.AnError
}

func TestJSONResponseEncoderError(t *testing.T) {
	rec := httptest.NewRecorder()
	err := http.JSONResponseEncoder(context.Background(), rec, marshalJSONError{})

	var marshalerError *json.MarshalerError
	assert.True(t, errors.As(err, &marshalerError))
	assert.Equal(t, assert.AnError, marshalerError.Err)
}

func TestWithStatusCode(t *testing.T) {
	type resp struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}

	response := http.WithStatusCode(resp{ID: "id", Text: "text"}, 202)

	statusCode, ok := response.(kithttp.StatusCoder)
	assert.True(t, ok)
	assert.Equal(t, 202, statusCode.StatusCode())

	body, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"id":"id","text":"text"}`, string(body))
}
