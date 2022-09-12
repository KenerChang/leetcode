package kclosestpointstoorigin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKClosest(t *testing.T) {
	cases := []struct {
		name   string
		wanted [][]int
		points [][]int
		k      int
	}{
		{"case 1", [][]int{{-2, 2}}, [][]int{{1, 3}, {-2, 2}}, 1},
		{"case 2", [][]int{{3, 3}, {-2, 4}}, [][]int{{3, 3}, {5, -1}, {-2, 4}}, 2},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.wanted, kClosest(c.points, c.k))
		})
	}
}
