package numberofislands

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	target := 1

	result := numIslands(grid)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestNumIslandsII(t *testing.T) {
	grid := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	target := 3

	result := numIslands(grid)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestNumIslandsIII(t *testing.T) {
	grid := [][]byte{}
	target := 0

	result := numIslands(grid)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestNumIslandsIV(t *testing.T) {
	grid := [][]byte{
		[]byte{'1', '0', '1'},
	}
	target := 2

	result := numIslands(grid)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestNumIslandsV(t *testing.T) {
	grid := [][]byte{
		[]byte{'0', '1', '1'},
	}
	target := 1

	result := numIslands(grid)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
