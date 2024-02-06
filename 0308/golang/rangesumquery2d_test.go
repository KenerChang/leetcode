package rangesumquery2d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumMatrix(t *testing.T) {
	numMatrix := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})

	assert.Equal(t, 8, numMatrix.SumRegion(2, 1, 4, 3))

	numMatrix.Update(3, 2, 2)

	assert.Equal(t, 10, numMatrix.SumRegion(2, 1, 4, 3))
}
