package longestconsecutivesequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{"100, 4, 200, 1, 3, 2", []int{100, 4, 200, 1, 3, 2}, 4},
		{"0, 3, 7, 2, 5, 8, 4, 6, 0, 1", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"100, 101, 102", []int{102, 100, 101}, 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, longestConsecutive(c.nums))
		})
	}
}
