package graphvalidtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidTree(t *testing.T) {
	assert.True(t, validTree(5, [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 4},
	},
	))

	assert.False(t, validTree(5, [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{1, 3},
		{1, 4},
	}))

	assert.True(t, validTree(1, [][]int{}))

	assert.True(t, validTree(2, [][]int{
		{1, 0},
	}))

	assert.True(t, validTree(3, [][]int{
		{1, 0},
		{2, 0},
	}))

	assert.True(t, validTree(3, [][]int{
		{2, 0},
		{2, 1},
	}))

	assert.False(t, validTree(5, [][]int{
		{0, 1},
		{0, 4},
		{1, 4},
		{2, 3},
	}))

	assert.False(t, validTree(6, [][]int{
		{0, 1},
		{0, 2},
		{3, 4},
		{3, 5},
	}))

}
