package missingranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	cases := []struct {
		nums   []int
		lower  int
		upper  int
		wanted [][]int
	}{
		{[]int{0, 1, 3, 50, 75}, 0, 99, [][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}}},
		{[]int{-1}, -1, -1, [][]int{}},
		{[]int{0}, -1, 1, [][]int{{-1, -1}, {1, 1}}},
		{[]int{}, -1, -1, [][]int{{-1, -1}}},
		{[]int{}, -1, 100, [][]int{{-1, 100}}},
		{[]int{0}, -10, 10, [][]int{{-10, -1}, {1, 10}}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("nums: %v, lower: %d, upper: %d", c.nums, c.lower, c.upper), func(t *testing.T) {
			assert.Equal(t, c.wanted, findMissingRanges(c.nums, c.lower, c.upper))
		})
	}
}
