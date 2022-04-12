package hindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndex(t *testing.T) {
	// assert.Equal(t, 3, hIndex([]int{3, 0, 6, 1, 5}))

	assert.Equal(t, 1, hIndex([]int{1, 3, 1}))

	assert.Equal(t, 0, hIndex([]int{0}))

	assert.Equal(t, 1, hIndex([]int{1}))

	assert.Equal(t, 4, hIndex([]int{5, 5, 5, 5, 1}))

	assert.Equal(t, 1, hIndex([]int{100}))

	assert.Equal(t, 2, hIndex([]int{100, 50}))
}
