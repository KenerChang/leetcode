package coinchange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	cases := []struct {
		wanted int
		coins  []int
		amount int
	}{
		{3, []int{1, 2, 5}, 11},
		{-1, []int{2}, 3},
		{0, []int{1}, 0},
		{2, []int{4, 7, 5}, 12},
		{3, []int{2, 3, 4}, 12},
		{20, []int{1, 2, 5}, 100},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v_%d", c.coins, c.amount), func(t *testing.T) {
			assert.Equal(t, c.wanted, coinChange(c.coins, c.amount))
		})
	}
}
