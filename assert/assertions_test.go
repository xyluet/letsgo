package assert_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	letsgoassert "github.com/xyluet/letsgo/assert"
)

func TestJSONLEq_EqualsJSONLString(t *testing.T) {
	assert.True(t, letsgoassert.JSONLEq(t,
		strings.Join([]string{`{"hello": "world"}`, `{"foo": "bar"}`}, "\n"),
		strings.Join([]string{`{"hello": "world"}`, `{"foo": "bar"}`}, "\n"),
	))
}
