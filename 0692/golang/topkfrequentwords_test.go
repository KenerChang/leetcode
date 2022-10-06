package topkfrequentwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		name   string
		wanted []string
		words  []string
		k      int
	}{
		{"case 1", []string{"i", "love"}, []string{"i", "love", "leetcode", "i", "love", "coding"}, 2},
		{"case 2", []string{"the", "is", "sunny", "day"}, []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4},
		{"case 3", []string{"i"}, []string{"i", "love", "leetcode", "i", "love", "coding"}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, topKFrequent(c.words, c.k))
		})
	}
}
