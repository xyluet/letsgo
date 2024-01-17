package assert

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// JSONLEq compares two JSONL strings and asserts that they are equal.
// It converts the JSONL strings to JSON format before performing the comparison.
func JSONLEq(t *testing.T, expected, actual string) bool {
	return assert.JSONEq(t, jsonlToJSON(t, expected), jsonlToJSON(t, actual))
}

func jsonlToJSON(t *testing.T, jsonl string) string {
	decoder := json.NewDecoder(bufio.NewReader(strings.NewReader(jsonl)))
	result := []any{}
	for {
		var data any
		err := decoder.Decode(&data)
		if err == io.EOF {
			break
		}

		assert.NoError(t, err)
		result = append(result, data)
	}
	encodedResult, err := json.Marshal(result)
	assert.NoError(t, err)
	return string(encodedResult)
}
