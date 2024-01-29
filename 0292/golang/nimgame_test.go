package nimgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWinNim(t *testing.T) {
	cases := []struct {
		n      int
		wanted bool
	}{
		{4, false},
		{1, true},
		{2, true},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("n: %d", c.n), func(t *testing.T) {
			assert.Equal(t, c.wanted, canWinNim(c.n))
		})
	}
}
