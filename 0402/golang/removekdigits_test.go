package removekdigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveKdigits(t *testing.T) {
	cases := []struct {
		wanted string
		num    string
		k      int
	}{
		{"1219", "1432219", 3},
		{"200", "10200", 1},
		{"0", "10", 2},
		{"56", "5678", 2},
		{"5", "5678", 3},
		{"0", "0", 1},
		{"0", "100", 2},
		{"0", "100", 1},
		{"99", "999", 1},
		{"9", "999", 2},
		{"0", "999", 3},
		{"0", "10001", 4},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("num:%s,k:%d", c.num, c.k), func(t *testing.T) {
			assert.Equal(t, c.wanted, removeKdigits(c.num, c.k))
		})
	}
}
