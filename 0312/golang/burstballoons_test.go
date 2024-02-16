package burstballoons

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCoins(t *testing.T) {
	cases := []struct {
		nums   []int
		wanted int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("nums: %v", c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, maxCoins(c.nums))
		})
	}
}
