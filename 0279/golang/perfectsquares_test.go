package perfectsquares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquares(t *testing.T) {
	assert.Equal(t, 3, numSquares(12))

	assert.Equal(t, 1, numSquares(4))

	assert.Equal(t, 1, numSquares(9))

	assert.Equal(t, 2, numSquares(13))

	assert.Equal(t, 3, numSquares(3))

	assert.Equal(t, 4, numSquares(15))

	assert.Equal(t, 2, numSquares(1000))

	assert.Equal(t, 1, numSquares(10000))

}
