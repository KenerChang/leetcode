package querykthsmallesttrimmednumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestTrimmedNumbers(t *testing.T) {
	cases := []struct {
		name    string
		wanted  []int
		nums    []string
		queries [][]int
	}{
		{"case 1", []int{2, 2, 1, 0}, []string{"102", "473", "251", "814"}, [][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}}},
		{"case 2", []int{3, 0}, []string{"24", "37", "96", "04"}, [][]int{{2, 1}, {2, 2}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, smallestTrimmedNumbers(c.nums, c.queries))
		})
	}
}
