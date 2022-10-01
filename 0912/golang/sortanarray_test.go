package sortanarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArray(t *testing.T) {
	cases := []struct {
		name   string
		wanted []int
		given  []int
	}{
		{"case 1", []int{1, 2, 3, 5}, []int{5, 2, 3, 1}},
		{"case 2", []int{0, 0, 1, 1, 2, 5}, []int{5, 1, 1, 2, 0, 0}},
		{"case 3", []int{1, 2, 3, 3, 4, 5, 6}, []int{1, 2, 3, 3, 4, 5, 6}},
		{"case 4", []int{1, 2, 3, 3, 4, 5, 6}, []int{6, 5, 4, 3, 3, 2, 1}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, sortArray(c.given))
		})
	}
}
