package maximalsquare

import "testing"

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	result := maximalSquare(matrix)

	target := 4
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaximalSquareII(t *testing.T) {
	matrix := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}

	result := maximalSquare(matrix)

	target := 1
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestMaximalSquareIII(t *testing.T) {
	matrix := [][]byte{
		{'0'},
	}

	result := maximalSquare(matrix)

	target := 0
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}
