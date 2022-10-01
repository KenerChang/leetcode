package rleiterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRLEIterat(t *testing.T) {
	iter := Constructor([]int{3, 8, 0, 9, 2, 5})

	assert.Equal(t, 8, iter.Next(2))
	assert.Equal(t, 8, iter.Next(1))
	assert.Equal(t, 5, iter.Next(1))
	assert.Equal(t, -1, iter.Next(2))
}
