package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameOfLifeI(t *testing.T) {
	board := [][]int{
		{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
	}

	gameOfLife(board)

	target := [][]int{
		{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0},
	}
	assert.Equal(t, target, board)
}

func TestGameOfLifeII(t *testing.T) {
	board := [][]int{
		{1, 0}, {1, 1},
	}

	gameOfLife(board)

	target := [][]int{
		{1, 1}, {1, 1},
	}
	assert.Equal(t, target, board)
}

func TestGameOfLifeIII(t *testing.T) {
	board := [][]int{
		{0},
	}

	gameOfLife(board)

	target := [][]int{
		{0},
	}
	assert.Equal(t, target, board)
}

func TestGameOfLifeIV(t *testing.T) {
	board := [][]int{
		{1},
	}

	gameOfLife(board)

	target := [][]int{
		{0},
	}
	assert.Equal(t, target, board)
}
