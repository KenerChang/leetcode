package stampingthesequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovesToStamp(t *testing.T) {
	cases := []struct {
		name   string
		wanted []int
		stamp  string
		target string
	}{
		{"case 1", []int{0, 2}, "abc", "ababc"},
		{"case 2", []int{3, 0, 1}, "abca", "aabcaca"},
		{"case 3", []int{0}, "a", "a"},
		{"case 4", []int{}, "abc", "cababc"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, movesToStamp(c.stamp, c.target))
		})
	}
}
