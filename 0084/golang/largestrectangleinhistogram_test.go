package largestrectangleinhistogram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	assert.Equal(t, 10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))

	assert.Equal(t, 4, largestRectangleArea([]int{2, 4}))

	assert.Equal(t, 8, largestRectangleArea([]int{2, 2, 3, 4}))

	assert.Equal(t, 9, largestRectangleArea([]int{2, 3, 3, 4}))

	assert.Equal(t, 6, largestRectangleArea([]int{3, 1, 1, 1, 2, 1}))

	assert.Equal(t, 8, largestRectangleArea([]int{3, 1, 1, 1, 2, 2, 2, 2}))

	assert.Equal(t, 10, largestRectangleArea([]int{3, 1, 1, 1, 2, 2, 2, 2, 2}))

	assert.Equal(t, 15, largestRectangleArea([]int{3, 1, 1, 15, 2, 2, 2, 2, 2}))

	assert.Equal(t, 18, largestRectangleArea([]int{4, 2, 2, 2, 3, 3, 3, 3, 3}))
}
