package pairsofsongswithtotaldurationsdivisibleby60

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumPairsDivisibleBy60(t *testing.T) {
	cases := []struct {
		wanted int
		time   []int
	}{
		{3, []int{30, 20, 150, 100, 40}},
		{3, []int{60, 60, 60}},
		{0, []int{50}},
		{0, []int{60}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.time), func(t *testing.T) {
			assert.Equal(t, c.wanted, numPairsDivisibleBy60(c.time))
		})
	}
}
