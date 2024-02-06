package kit_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/stretchr/testify/assert"
	"github.com/xyluet/letsgo/exp/kit"
)

func TestEndpoint(t *testing.T) {
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
	anDecodeRequestFunc := func(ctx context.Context, req *http.Request) (requestType, error) {
		defer req.Body.Close()
		b, _ := io.ReadAll(req.Body)
		var rt requestType
		_ = json.Unmarshal(b, &rt)

		return rt, nil
	}
	anEncodeResponseFunc := func(ctx context.Context, w http.ResponseWriter, resp responseType) error {
		return json.NewEncoder(w).Encode(resp)
	}

	generic := kit.NewGeneric[requestType, responseType]()

	handler := kithttp.NewServer(
		generic.Endpoint(anEndpoint),
		generic.DecodeRequestFunc(anDecodeRequestFunc),
		generic.EncodeResponseFunc(anEncodeResponseFunc),
	)
	w, r := httptest.NewRecorder(), httptest.NewRequest("GET", "/", strings.NewReader(`{"name":"a name"}`))
	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"name":"a name","what":"foo"}`, w.Body.String())
}
