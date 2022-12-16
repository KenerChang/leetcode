package sentencescreenfitting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordsTyping(t *testing.T) {
	cases := []struct {
		wanted   int
		sentence []string
		row      int
		col      int
	}{
		{1, []string{"hello", "world"}, 2, 8},
		{2, []string{"a", "bcd", "e"}, 3, 6},
		{2, []string{"a", "bcd", "e"}, 4, 6},
		{1, []string{"i", "had", "apple", "pie"}, 4, 5},
		{0, []string{"hello", "world"}, 1, 5},
		{1, []string{"hello", "world"}, 2, 5},
		{1, []string{"hello", "world"}, 3, 5},
		{0, []string{"hello", "worldiiii"}, 2, 5},
		{10, []string{"f", "p", "a"}, 8, 7},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("sentence:%v,row:%d,col:%d", c.sentence, c.row, c.col), func(t *testing.T) {
			assert.Equal(t, c.wanted, wordsTyping(c.sentence, c.row, c.col))
		})
	}
}
