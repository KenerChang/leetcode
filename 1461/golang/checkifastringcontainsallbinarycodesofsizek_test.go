package checkifastringcontainsallbinarycodesofsizek

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasAllCodes(t *testing.T) {
	cases := []struct {
		name   string
		wanted bool
		s      string
		k      int
	}{
		{"00110110", true, "00110110", 2},
		{"0110", true, "0110", 1},
		{"0110", false, "0110", 2},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, hasAllCodes(c.s, c.k))
		})
	}
}
