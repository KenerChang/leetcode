package mirrorreflection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMirrorReflection(t *testing.T) {
	cases := []struct {
		wanted int
		p      int
		q      int
	}{
		{2, 2, 1},
		{1, 3, 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("p: %d, q: %d", c.p, c.q), func(t *testing.T) {
			assert.Equal(t, c.wanted, mirrorReflection(c.p, c.q))
		})
	}
}
