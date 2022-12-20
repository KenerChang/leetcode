package dietplanperformance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDietPlanPerformance(t *testing.T) {
	cases := []struct {
		wanted   int
		calories []int
		k        int
		lower    int
		upper    int
	}{
		{0, []int{1, 2, 3, 4, 5}, 1, 3, 3},
		{-5, []int{1, 2, 3, 4, 5}, 1, 6, 6},
		{5, []int{1, 2, 3, 4, 5}, 1, 0, 0},
		{1, []int{3, 2}, 2, 0, 1},
		{0, []int{6, 5, 0, 0}, 2, 1, 5},
		{1, []int{6, 5, 1}, 3, 1, 5},
		{-1, []int{6, 5, 1}, 3, 15, 15},
		{0, []int{6, 5, 1}, 3, 8, 15},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("calories:%v,k:%d,lower:%d,upper:%d", c.calories, c.k, c.lower, c.upper), func(t *testing.T) {
			assert.Equal(t, c.wanted, dietPlanPerformance(c.calories, c.k, c.lower, c.upper))
		})
	}
}
