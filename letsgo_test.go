package letsgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xyluet/letsgo"
)

func TestMust(t *testing.T) {
	v := letsgo.Must[int](0, nil)
	assert.Equal(t, 0, v)
	assert.Panics(t, func() { letsgo.Must[int](0, errors.New("!")) })
}
