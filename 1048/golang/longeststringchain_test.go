package longeststringchain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestStrChain(t *testing.T) {
	cases := []struct {
		wanted int
		words  []string
	}{
		{4, []string{"a", "b", "ba", "bca", "bda", "bdca"}},
		{5, []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}},
		{1, []string{"a", "b"}},
		{2, []string{"ac", "b", "a"}},
		{3, []string{"a", "b", "ac", "bd", "bde"}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.words), func(t *testing.T) {
			assert.Equal(t, c.wanted, longestStrChain(c.words))
		})
	}
}
