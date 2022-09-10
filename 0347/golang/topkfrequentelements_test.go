package topkfrequentelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		name   string
		wanted []int
		nums   []int
		k      int
	}{
		{"test case 1", []int{1, 2}, []int{1, 1, 1, 2, 2, 3}, 2},
		{"test case 2", []int{1}, []int{1}, 1},
		{"test case 3", []int{1, 2, 3}, []int{1, 1, 1, 1, 2, 2, 3, 3, 4}, 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, topKFrequent(c.nums, c.k))
		})
	}
}
