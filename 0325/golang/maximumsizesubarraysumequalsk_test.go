package maximumsizesubarraysumequalsk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArrayLen(t *testing.T) {
	cases := []struct {
		wanted int
		nums   []int
		k      int
	}{
		{4, []int{1, -1, 5, -2, 3}, 3},
		{2, []int{-2, -1, 2, 1}, 1},
		{0, []int{1, -1}, 2},
		{3, []int{1, 2, 3}, 6},
		{1, []int{1}, 1},
		{0, []int{1}, 2},
		{3, []int{1, 2, 3, -1, -2, -3}, 6},
		{6, []int{1, 2, 3, -3, -2, -1}, 0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("nums:%v_k:%d", c.nums, c.k), func(t *testing.T) {
			assert.Equal(t, c.wanted, maxSubArrayLen(c.nums, c.k))
		})
	}
}
