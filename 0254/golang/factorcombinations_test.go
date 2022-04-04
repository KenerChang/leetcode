package factorcombinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFactors(t *testing.T) {
	target := [][]int{}

	assert.ElementsMatch(t, target, getFactors(1))

	target = [][]int{
		{2, 6},
		{3, 4},
		{2, 2, 3},
	}

	assert.ElementsMatch(t, target, getFactors(12))

	target = [][]int{
		{2, 37},
	}

	assert.ElementsMatch(t, target, getFactors(74))

	target = [][]int{}

	assert.ElementsMatch(t, target, getFactors(37))

	target = [][]int{
		{2, 16}, {4, 8}, {2, 2, 8}, {2, 4, 4}, {2, 2, 2, 4}, {2, 2, 2, 2, 2},
	}

	assert.ElementsMatch(t, target, getFactors(32))
}
