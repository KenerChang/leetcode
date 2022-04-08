package missingnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	assert.Equal(t, 2, missingNumber([]int{3, 0, 1}))

	assert.Equal(t, 2, missingNumber([]int{0, 1}))

	assert.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

	assert.Equal(t, 1, missingNumber([]int{0}))
}
