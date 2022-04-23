package maxpointsonaline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPoints(t *testing.T) {
	assert.Equal(t, 3, maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))

	assert.Equal(t, 4, maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))

	assert.Equal(t, 4, maxPoints([][]int{{1, 1}, {2, 2}, {2, 1}, {1, 2}, {3, 1}, {3, 2}, {4, 2}}))

	assert.Equal(t, 1, maxPoints([][]int{{1, 1}}))

	assert.Equal(t, 2, maxPoints([][]int{{1, 1}, {2, 2}}))
}
