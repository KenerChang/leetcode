package subarraysumequalsk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	cases := []struct {
		wanted int
		nums   []int
		k      int
	}{
		{2, []int{1, 1, 1}, 2},
		{2, []int{1, 2, 3}, 3},
		{1, []int{1, 3, 2}, 3},
		{3, []int{1, 1, 1}, 1},
		{0, []int{1, 1, 1}, 0},
		{2, []int{1, -1, 1}, 0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, subarraySum(c.nums, c.k))
		})
	}
}
