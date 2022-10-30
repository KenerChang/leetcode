package longestincreasingsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	cases := []struct {
		wanted int
		nums   []int
	}{
		{4, []int{10, 9, 2, 5, 3, 7, 101, 18}},
		{4, []int{0, 1, 0, 3, 2, 3}},
		{1, []int{7, 7, 7, 7, 7}},
		{3, []int{1, 2, 3}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, lengthOfLIS(c.nums))
		})
	}
}
