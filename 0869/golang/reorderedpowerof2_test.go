package reorderedpowerof2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderedPowerOf2(t *testing.T) {
	cases := []struct {
		name string
		want bool
		n    int
	}{
		{"test case 1", true, 1},
		{"test case 2", false, 10},
		{"test case 3", true, 125},
		{"test case 4", true, 4021},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, reorderedPowerOf2(c.n))
		})
	}
}
