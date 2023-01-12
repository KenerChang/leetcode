package palindromenumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		wanted bool
		x      int
	}{
		{true, 121},
		{false, -121},
		{false, 10},
		{true, 0},
		{true, 1001},
		{false, 51},
		{true, 5},
		{true, 55},
		{true, 555},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.x), func(t *testing.T) {
			assert.Equal(t, c.wanted, isPalindrome(c.x))
		})
	}
}
