package nextclosesttime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextClosestTime(t *testing.T) {
	cases := []struct {
		wanted string
		time   string
	}{
		{"19:39", "19:34"},
		{"22:22", "23:59"},
		{"00:00", "00:00"},
		{"11:11", "11:11"},
		{"11:21", "11:12"},
		{"23:02", "23:00"},
	}

	for _, c := range cases {
		t.Run(c.time, func(t *testing.T) {
			assert.Equal(t, c.wanted, nextClosestTime(c.time))
		})
	}
}
