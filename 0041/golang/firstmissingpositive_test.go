package firstmissingpositive

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	cases := []struct {
		nums   []int
		wanted int
	}{
		{
			nums:   []int{1, 2, 0},
			wanted: 3,
		},
		{
			nums:   []int{3, 4, -1, 1},
			wanted: 2,
		},
		{
			nums:   []int{7, 8, 9, 11, 12},
			wanted: 1,
		},
		{
			nums:   []int{3, 5, 1},
			wanted: 2,
		},
		{
			nums:   []int{-21, -11, -1},
			wanted: 1,
		},
		{
			nums:   []int{2, 1},
			wanted: 3,
		},
		{
			nums:   []int{5, 3, 2, 1},
			wanted: 4,
		},
		{
			nums:   []int{-1, 4, 2, 1, 9, 10},
			wanted: 3,
		},
		{
			nums:   []int{98, 93, 95, 10, 91, 4, 90, 88, 56, 84, 65, 62, 83, 80, 78, 60, 73, 77, 76, 29, 63, 12, 57, 17, 69, 68, 50, 11, 31, 33, 8, 42, 38, 7, 0, 37, 48, 26, 20, 44, 46, 43, 52, 51, 47, 18, 49, 58, 2, 39, 30, 81, 22, 55, 36, 40, 15, 27, 21, 32, 64, 41, 53, 19, 28, 24, 9, 25, 3, 59, 66, 82, 61, 70, 23, 34, 71, 54, 74, -1, 1, 45, 14, 79, 5, 35, 13, 72, 75, 85, 87, 6, 16, 86, 67, 89, 94, 92, 96, 97},
			wanted: 99,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprint(c.nums), func(t *testing.T) {
			assert.Equal(t, c.wanted, firstMissingPositive(c.nums))
		})
	}
}
