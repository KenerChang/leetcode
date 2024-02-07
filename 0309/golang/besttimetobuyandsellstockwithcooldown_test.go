package besttimetobuyandsellstockwithcooldown

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		wanted int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		// {[]int{1}, 0},
		// {[]int{5, 4, 3, 2, 1}, 0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("prices: %v", c.prices), func(t *testing.T) {
			assert.Equal(t, c.wanted, maxProfit(c.prices))
		})
	}
}
