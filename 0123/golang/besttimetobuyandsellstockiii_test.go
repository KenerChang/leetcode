package besttimetobuyandsellstockiii

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
		{
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			wanted: 6,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			wanted: 4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			wanted: 0,
		},
		{
			prices: []int{1, 4, 7, 0, 2, 4, 0, 1, 2},
			wanted: 10,
		},
		{
			prices: []int{1},
			wanted: 0,
		},
		{
			prices: []int{1, 3},
			wanted: 2,
		},
		{
			prices: []int{3, 1},
			wanted: 0,
		},
		{
			prices: []int{1, 3, 2},
			wanted: 2,
		},
		{
			prices: []int{1, 4, 2, 7},
			wanted: 8,
		},
		{
			prices: []int{3, 2, 6, 5, 0, 3},
			wanted: 7,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("prices: %v", c.prices), func(t *testing.T) {
			assert.Equal(t, c.wanted, maxProfit(c.prices))
		})
	}
}
