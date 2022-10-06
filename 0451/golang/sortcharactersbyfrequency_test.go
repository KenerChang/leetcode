package sortcharactersbyfrequency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	cases := []struct {
		name   string
		wanted string
		s      string
	}{
		{"", "eert", "tree"},
		{"", "aaaccc", "cccaaa"},
		{"", "bbAa", "Aabb"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, frequencySort(c.s))
		})
	}
}
