package maximumlengthofaconcatenatedstringwithuniquecharacters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLength(t *testing.T) {
	cases := []struct {
		wanted int
		arr    []string
	}{
		{4, []string{"un", "iq", "ue"}},
		{6, []string{"cha", "r", "act", "ers"}},
		{26, []string{"abcdefghijklmnopqrstuvwxyz"}},
		{3, []string{"a", "b", "c"}},
		{3, []string{"a", "ab", "abc"}},
		{0, []string{"aa", "bb"}},
		{2, []string{"aa", "bc"}},
		{11, []string{"ab", "cd", "cde", "cdef", "efg", "fgh", "abxyz"}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.arr), func(t *testing.T) {
			assert.Equal(t, c.wanted, maxLength(c.arr))
		})
	}
}
