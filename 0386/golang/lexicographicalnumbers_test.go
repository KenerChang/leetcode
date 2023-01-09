package lexicographicalnumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexicalOrder(t *testing.T) {
	cases := []struct {
		wanted []int
		n      int
	}{
		{[]int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}, 13},
		{[]int{1, 2}, 2},
		{[]int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 3, 4, 5, 6, 7, 8, 9}, 20},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.n), func(t *testing.T) {
			assert.Equal(t, c.wanted, lexicalOrder(c.n))
		})
	}
}
