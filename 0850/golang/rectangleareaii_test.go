package rectangleareaii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangleArea(t *testing.T) {
	cases := []struct {
		name   string
		wanted int
		given  [][]int
	}{
		{"case 1", 6, [][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}},
		{"case 2", 49, [][]int{{0, 0, 1000000000, 1000000000}}},
		{"case 3", 18, [][]int{{0, 0, 3, 3}, {2, 0, 5, 3}, {1, 1, 4, 4}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, rectangleArea(c.given))
		})
	}
}
