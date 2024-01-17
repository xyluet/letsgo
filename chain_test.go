package letsgo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xyluet/letsgo"
)

func addByOne(input int) int {
	return input + 1
}

func multiplyByTwo(input int) int {
	return input * 2
}

func TestChain(t *testing.T) {
	is := assert.New(t)

	chain := letsgo.Chain(addByOne, multiplyByTwo)
	output := chain(5)
	is.Equal(12, output)
}

func TestChainNoMiddlewares(t *testing.T) {
	is := assert.New(t)

	chain := letsgo.Chain(addByOne)
	output := chain(5)
	is.Equal(6, output)
}

func TestChainComposed(t *testing.T) {
	is := assert.New(t)

	chained := letsgo.Chain(addByOne, multiplyByTwo)
	chainedWithSquare := letsgo.Chain(chained, func(input int) int {
		return input * input
	})

	output := chainedWithSquare(5)
	is.Equal(144, output)
}
