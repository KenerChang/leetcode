package findminimuminrotatedsortedarrayii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, findMin([]int{1, 3, 5}))

	assert.Equal(t, 0, findMin([]int{2, 2, 2, 0, 1}))

	assert.Equal(t, 0, findMin([]int{2, 2, 2, 1, 0}))

	assert.Equal(t, 2, findMin([]int{2}))

	assert.Equal(t, 5, findMin([]int{5, 5, 5}))

	assert.Equal(t, 0, findMin([]int{1, 2, 2, 2, 0}))

	assert.Equal(t, 0, findMin([]int{1, 2, 3, 0, 1, 1, 1}))

	assert.Equal(t, 0, findMin([]int{1, 2, 3, 0, 1, 1, 1, 1, 1, 1, 1}))

	assert.Equal(t, 0, findMin([]int{1, 1, 1, 1, 1, 1, 2, 3, 0, 1, 1}))

	assert.Equal(t, -5, findMin([]int{2, 3, 4, 5, 6, -5, -4, -3, -2, -1, 0}))

	assert.Equal(t, -5, findMin([]int{1, 2, 3, 4, 5, 6, -5, -4, -3, -2, -1}))
}
