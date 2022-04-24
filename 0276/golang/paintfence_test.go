package paintfence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWays(t *testing.T) {
	assert.Equal(t, 6, numWays(3, 2))

	assert.Equal(t, 1, numWays(1, 1))

	assert.Equal(t, 42, numWays(7, 2))

	assert.Equal(t, 27408, numWays(10, 3))
}
