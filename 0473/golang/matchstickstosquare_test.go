package matchstickstosquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakesquare(t *testing.T) {
	cases := []struct {
		name        string
		want        bool
		matchsticks []int
	}{
		{"test case 1", true, []int{1, 1, 2, 2, 2}},
		{"test case 2", false, []int{3, 3, 3, 3, 4}},
		{"test case 3", true, []int{1, 2, 1, 2, 1, 2, 1, 2}},
		{"test case 4", true, []int{1, 3, 3, 3, 1, 1}},
		{"test case 5", false, []int{1, 2, 1, 2, 1, 2}},
		{"test case 6", true, []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}},
		{"test case 7", true, []int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, makesquare(c.matchsticks))
		})
	}
}
