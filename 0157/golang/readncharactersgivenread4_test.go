package readncharactersgivenread4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name   string
		wanted string
		file   string
		n      int
	}{
		{"abc", "abc", "abc", 4},
		{"abcde", "abcde", "abcde", 5},
		{"abcdABCD1234", "abcdABCD1234", "abcdABCD1234", 12},
		{"abcdABCD1234", "a", "abcdABCD1234", 1},
		{"abc", "abc", "abc", 20},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := make([]byte, min(c.n, len(c.file)))
			read4 := read4Impl(c.file)
			readFun := solution(read4)

			read := readFun(buf, c.n)
			assert.Equal(t, c.wanted, string(buf))
			assert.Equal(t, min(c.n, len(c.file)), read)
		})
	}
}
