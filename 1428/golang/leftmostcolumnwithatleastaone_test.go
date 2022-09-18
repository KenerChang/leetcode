package leftmostcolumnwithatleastaone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftMostColumnWithOne(t *testing.T) {
	cases := []struct {
		name   string
		wanted int
		given  [][]int
	}{
		{"case 1", 0, [][]int{{0, 0}, {1, 1}}},
		{"case 2", 1, [][]int{{0, 0}, {0, 1}}},
		{"case 3", -1, [][]int{{0, 0}, {0, 0}}},
		{"case 4", 0, [][]int{{1, 1, 1}, {0, 1, 1}, {1, 1, 1}}},
		{"case 5", 1, [][]int{{0, 1}, {0, 0}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			matrix := BinaryMatrix{
				Get:        GetFactory(c.given),
				Dimensions: GetDimensionsFactory(c.given),
			}

			assert.Equal(t, c.wanted, leftMostColumnWithOne(matrix))
		})
	}
}
