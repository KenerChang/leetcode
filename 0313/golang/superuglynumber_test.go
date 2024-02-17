package superuglynumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthSuperUglyNumber(t *testing.T) {
	cases := []struct {
		n      int
		primes []int
		wanted int
	}{
		{2, []int{2, 7, 13, 19}, 2},
		{12, []int{2, 7, 13, 19}, 32},
		{1, []int{2, 3, 5}, 1},
		// {100000, []int{2, 7, 13, 19}, 32},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("n: %d, primes: %v", c.n, c.primes), func(t *testing.T) {
			assert.Equal(t, c.wanted, nthSuperUglyNumber(c.n, c.primes))
		})
	}
}
