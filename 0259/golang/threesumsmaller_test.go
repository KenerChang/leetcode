package threesumsmaller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumSmaller(t *testing.T) {
	// assert.Equal(t, 2, threeSumSmaller([]int{-2, 0, 1, 3}, 2))

	// assert.Equal(t, 0, threeSumSmaller([]int{}, 0))

	// assert.Equal(t, 965, threeSumSmaller([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9}, 22))

	assert.Equal(t, 1, threeSumSmaller([]int{-1, 1, -1, -1}, -1))
}
