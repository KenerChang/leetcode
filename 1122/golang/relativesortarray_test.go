package relativesortarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeSortArray(t *testing.T) {
	cases := []struct {
		name   string
		wanted []int
		arr1   []int
		arr2   []int
	}{
		{"case 1", []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}, []int{2, 3, 1, 3, 2, 4, 6, 19, 2, 9, 7}, []int{2, 1, 4, 3, 9, 6}},
		{"case 2", []int{22, 28, 8, 6, 17, 44}, []int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, relativeSortArray(c.arr1, c.arr2))
		})
	}
}
