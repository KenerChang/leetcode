package flipgameii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWin(t *testing.T) {
	cases := []struct {
		name   string
		wanted bool
		given  string
	}{
		{"case 1", true, "++++"},
		{"case 2", false, "+"},
		{"case 3", true, "++"},
		{"case 4", false, "--"},
		{"case 5", true, "++++++"},
		{"case 6", false, "+++++"},
		{"case 7", false, "+--+"},
		{"case 8", false, "+++++++++"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, canWin(c.given))
		})
	}
}
