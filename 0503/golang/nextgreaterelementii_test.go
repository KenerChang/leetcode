package nextgreaterelementii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElements(t *testing.T) {
	cases := []struct {
		wanted []int
		nums   []int
	}{
		{[]int{2, -1, 2}, []int{1, 2, 1}},
		{[]int{2, 3, 4, -1, 4}, []int{1, 2, 3, 4, 3}},
		{[]int{-1, 3, 3}, []int{3, 2, 1}},
		{[]int{4, -1, 2}, []int{2, 4, 1}},
		{[]int{-1}, []int{1}},
		{[]int{-1, -1, -1, -1, -1}, []int{1, 1, 1, 1, 1}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, nextGreaterElements(c.nums))
		})
	}

}
