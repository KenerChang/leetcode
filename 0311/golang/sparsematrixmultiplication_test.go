package sparsematrixmultiplication

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	cases := []struct {
		mat1   [][]int
		mat2   [][]int
		wanted [][]int
	}{
		{[][]int{{1, 0, 0}, {-1, 0, 3}}, [][]int{{7, 0, 0}, {0, 0, 0}, {0, 0, 1}}, [][]int{{7, 0, 0}, {-7, 0, 3}}},
		{[][]int{{0}}, [][]int{{0}}, [][]int{{0}}},
		{[][]int{{1, -5}}, [][]int{{12}, {-1}}, [][]int{{17}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("mat1: %v, mat2: %v", c.mat1, c.mat2), func(t *testing.T) {
			assert.Equal(t, c.wanted, multiply(c.mat1, c.mat2))
		})
	}
}
