package strobogrammaticnumberiii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWinNim(t *testing.T) {
	cases := []struct {
		low    string
		high   string
		wanted int
	}{
		{"50", "100", 3},
		{"0", "0", 1},
		{"8", "8", 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("low: %s, hight: %s", c.low, c.high), func(t *testing.T) {
			assert.Equal(t, c.wanted, strobogrammaticInRange(c.low, c.high))
		})
	}
}
