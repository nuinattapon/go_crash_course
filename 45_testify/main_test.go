package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert.EqualValues(t, Calculate(2), 4)
	// require.EqualValues(t, Calculate(2), 4)
}
