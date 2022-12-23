package maxareaofisland

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAreaOfIslandI(t *testing.T) {
	grids := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	assert.Equal(t, 6, maxAreaOfIsland(grids))
}

func TestMaxAreaOfIslandII(t *testing.T) {
	grids := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	assert.Equal(t, 0, maxAreaOfIsland(grids))
}

func TestMaxAreaOfIslandIII(t *testing.T) {
	grids := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	assert.Equal(t, 1, maxAreaOfIsland(grids))
}

func TestMaxAreaOfIslandIV(t *testing.T) {
	grids := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	assert.Equal(t, 2, maxAreaOfIsland(grids))
}

func TestMaxAreaOfIslandV(t *testing.T) {
	grids := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
	}

	assert.Equal(t, 1, maxAreaOfIsland(grids))
}
