package nqueens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	// solveNQueens(4)
	target := [][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}
	assert.ElementsMatch(t, target, solveNQueens(4))

	assert.ElementsMatch(t, [][]string{{"Q"}}, solveNQueens(1))
}
