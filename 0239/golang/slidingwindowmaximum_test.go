package slidingwindowmaximum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	cases := []struct {
		name   string
		wanted []int
		nums   []int
		k      int
	}{
		{"case 1", []int{3, 3, 5, 5, 6, 7}, []int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
		{"case 2", []int{1}, []int{1}, 1},
		{"case 3", []int{1, -1}, []int{1, -1}, 1},
		{"case 4", []int{7, 4}, []int{7, 2, 4}, 2},
		{"case 5", []int{5, 3, 4}, []int{5, 3, 4}, 1},
		{"case 6", []int{3, 3, 2, 5}, []int{1, 3, 1, 2, 0, 5}, 3},
		{"case 7", []int{10, 10, 9, 2}, []int{9, 10, 9, -7, -4, -8, 2, -6}, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, maxSlidingWindow(c.nums, c.k))
		})
	}
}
