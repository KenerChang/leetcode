package reversepairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePairs(t *testing.T) {
	cases := []struct {
		name   string
		wanted int
		nums   []int
	}{
		{"case 1", 2, []int{1, 3, 2, 3, 1}},
		{"case 2", 3, []int{2, 4, 3, 5, 1}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, reversePairs(c.nums))
		})
	}
}
