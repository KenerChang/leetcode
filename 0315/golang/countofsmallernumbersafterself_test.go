package countofsmallernumbersafterself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmaller(t *testing.T) {
	cases := []struct {
		name   string
		wanted []int
		given  []int
	}{
		{"case 1", []int{2, 1, 1, 0}, []int{5, 2, 6, 1}},
		{"case 2", []int{0}, []int{-1}},
		{"case 3", []int{0, 0}, []int{-1, -1}},
		{"case 4", []int{0, 0, 0, 0, 0}, []int{1, 2, 3, 4, 5}},
		{"case 5", []int{4, 3, 2, 1, 0}, []int{5, 4, 3, 2, 1}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, countSmaller(c.given))
		})
	}
}
