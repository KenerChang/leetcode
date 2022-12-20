package maximumaveragesubarrayi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	cases := []struct {
		wanted float64
		nums   []int
		k      int
	}{
		{float64(12.75), []int{1, 12, -5, -6, 50, 3}, 4},
		{float64(5), []int{5}, 1},
		{float64(1), []int{1, 1, 1}, 3},
		{float64(247), []int{1, 2, 3, 4, -4, -4, -4, 1000}, 4},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("nums:%v,k:%d", c.nums, c.k), func(t *testing.T) {
			assert.Equal(t, c.wanted, findMaxAverage(c.nums, c.k))
		})
	}
}
