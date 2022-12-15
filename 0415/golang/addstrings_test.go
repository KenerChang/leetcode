package addstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStrings(t *testing.T) {
	cases := []struct {
		wanted string
		num1   string
		num2   string
	}{
		{"134", "11", "123"},
		{"533", "456", "77"},
		{"100", "100", "0"},
		{"0", "0", "0"},
		{"1000", "999", "1"},
		{"1000", "1", "999"},
		{"1000", "500", "500"},
		{"1000", "501", "499"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("num1:%s_num2:%s", c.num1, c.num2), func(t *testing.T) {
			assert.Equal(t, c.wanted, addStrings(c.num1, c.num2))
		})
	}
}
