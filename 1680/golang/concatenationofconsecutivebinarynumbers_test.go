package concatenationofconsecutivebinarynumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatenatedBinary(t *testing.T) {
	cases := []struct {
		wanted int
		n      int
	}{
		{1, 1},
		{6, 2},
		{27, 3},
		{505379714, 12},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.n), func(t *testing.T) {
			assert.Equal(t, c.wanted, concatenatedBinary(c.n))
		})
	}
}
