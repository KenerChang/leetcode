package movezeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	cases := []struct {
		nums   []int
		wanted []int
	}{
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, -1}, []int{-1, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("nums: %d", c.nums), func(t *testing.T) {
			moveZeroes(c.nums)
			assert.Equal(t, c.nums, c.wanted)
		})
	}
}
