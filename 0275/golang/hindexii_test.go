package hindexii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndex(t *testing.T) {
	assert.Equal(t, 3, hIndex([]int{0, 1, 3, 5, 6}))

	assert.Equal(t, 2, hIndex([]int{1, 2, 100}))

	assert.Equal(t, 5, hIndex([]int{8, 8, 8, 8, 8}))

	assert.Equal(t, 4, hIndex([]int{4, 4, 4, 4, 4}))

	assert.Equal(t, 4, hIndex([]int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4}))

	assert.Equal(t, 0, hIndex([]int{0, 0, 0}))

	assert.Equal(t, 1, hIndex([]int{0, 0, 1}))

	assert.Equal(t, 2, hIndex([]int{1, 2, 2, 2}))

	assert.Equal(t, 3, hIndex([]int{1, 4, 7, 9}))

	assert.Equal(t, 1, hIndex([]int{1, 1, 1, 1}))

	assert.Equal(t, 1, hIndex([]int{0, 0, 1, 1}))
}
