package editdistance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	cases := []struct {
		wanted int
		word1  string
		word2  string
	}{
		{3, "horse", "ros"},
		{5, "intention", "execution"},
		{1, "a", "b"},
		{1, "a", ""},
		{1, "ab", "ac"},
		{2, "aa", "bb"},
		{3, "aaa", "bbb"},
		{0, "a", "a"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s_%s", c.word1, c.word2), func(t *testing.T) {
			assert.Equal(t, c.wanted, minDistance(c.word1, c.word2))
		})
	}
}
