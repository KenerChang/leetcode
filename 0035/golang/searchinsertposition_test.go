package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))

	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))

	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))

	assert.Equal(t, 0, searchInsert([]int{1, 3, 5, 6}, 0))

	assert.Equal(t, 0, searchInsert([]int{}, 7))

	assert.Equal(t, 0, searchInsert([]int{8}, 7))

	assert.Equal(t, 0, searchInsert([]int{7}, 7))

	assert.Equal(t, 1, searchInsert([]int{5}, 7))
}
