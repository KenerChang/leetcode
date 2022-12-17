package ransomnote

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	cases := []struct {
		wanted     bool
		ransomNote string
		magazine   string
	}{
		{false, "a", "b"},
		{false, "aa", "ab"},
		{true, "aa", "aab"},
		{false, "aa", "a"},
		{true, "aab", "aab"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("ransomNote:%s,magazine:%s", c.ransomNote, c.magazine), func(t *testing.T) {
			assert.Equal(t, c.wanted, canConstruct(c.ransomNote, c.magazine))
		})
	}
}
