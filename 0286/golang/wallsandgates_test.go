package wallsandgates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	cases := []struct {
		rooms  [][]int
		wanted [][]int
	}{
		{[][]int{{-1}}, [][]int{{-1}}},
		{
			[][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}},
			[][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("rooms: %v", c.rooms), func(t *testing.T) {
			wallsAndGates(c.rooms)
			assert.Equal(t, c.rooms, c.wanted)
		})
	}
}
