package smallestrectangleenclosingblackpixels

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinArea(t *testing.T) {
	cases := []struct {
		image  [][]byte
		x      int
		y      int
		wanted int
	}{
		{[][]byte{{'0', '0', '1', '0'}, {'0', '1', '1', '0'}, {'0', '1', '0', '0'}}, 0, 2, 6},
		{[][]byte{{'1'}}, 0, 0, 1},
		{[][]byte{{'0', '1', '0', '0'}, {'0', '1', '1', '0'}, {'0', '1', '0', '0'}}, 0, 1, 6},
		{[][]byte{{'0', '0', '1', '1'}, {'0', '1', '1', '0'}, {'0', '1', '0', '0'}}, 0, 1, 9},
		{[][]byte{{'0', '0', '1', '0'}, {'0', '0', '1', '0'}, {'0', '0', '0', '0'}}, 0, 2, 2},
		{[][]byte{{'1', '1', '1', '1'}}, 0, 2, 4},
		{[][]byte{{'1'}, {'1'}, {'1'}, {'1'}}, 2, 0, 4},
		{[][]byte{{'1', '1', '1', '1'}, {'1', '1', '1', '1'}, {'1', '1', '1', '1'}, {'1', '1', '1', '1'}}, 0, 2, 16},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("image: %v, x: %d, y: %d", c.image, c.x, c.y), func(t *testing.T) {
			assert.Equal(t, c.wanted, minArea(c.image, c.x, c.y))
		})
	}
}
