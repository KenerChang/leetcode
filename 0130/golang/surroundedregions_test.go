package surroundedregions

import (
	"testing"
)

func verify(target, result [][]byte) bool {
	if len(target) != len(result) {
		return false
	}

	for i, row1 := range target {
		if len(row1) != len(result[i]) {
			return false
		}

		for j, cell1 := range row1 {
			cell2 := result[i][j]

			if cell1 != cell2 {
				return false
			}
		}
	}
	return true
}

func TestSolve(t *testing.T) {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'X'},
		[]byte{'X', 'X', 'O', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
	}

	target := [][]byte{
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
	}

	solve(board)

	if !verify(target, board) {
		t.Errorf("expect: %+v, got %+v", target, board)
	}
}

func TestSolveII(t *testing.T) {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'O'},
		[]byte{'X', 'X', 'O', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
	}

	target := [][]byte{
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'O'},
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
	}

	solve(board)

	if !verify(target, board) {
		t.Errorf("expect: %+v, got %+v", target, board)
	}
}

func TestSolveIII(t *testing.T) {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'O'},
		[]byte{'X', 'X', 'O', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
	}

	target := [][]byte{
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'O'},
		[]byte{'X', 'X', 'O', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
	}

	solve(board)

	if !verify(target, board) {
		t.Errorf("expect: %+v, got %+v", target, board)
	}
}

func TestSolveIV(t *testing.T) {
	board := [][]byte{
		[]byte{'X', 'O', 'X', 'O', 'X', 'O'},
		[]byte{'O', 'X', 'O', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'O', 'X', 'O'},
		[]byte{'O', 'X', 'O', 'X', 'O', 'X'},
	}

	target := [][]byte{
		[]byte{'X', 'O', 'X', 'O', 'X', 'O'},
		[]byte{'O', 'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'X', 'X', 'X', 'O'},
		[]byte{'O', 'X', 'O', 'X', 'O', 'X'},
	}

	solve(board)

	if !verify(target, board) {
		t.Errorf("expect: %+v, got %+v", target, board)
	}
}
