package removecoloredpiecesifbothneighborsarethesamecolor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWinnerOfGame(t *testing.T) {
	cases := []struct {
		name   string
		wanted bool
		given  string
	}{
		{"AAABABB", true, "AAABABB"},
		{"AA", false, "AA"},
		{"ABBBBBBBAAA", false, "ABBBBBBBAAA"},
		{"BB", false, "BB"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, winnerOfGame(c.given))
		})
	}
}
