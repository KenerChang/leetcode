package removeinvalidparentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveInvalidParentheses(t *testing.T) {
	cases := []struct {
		s      string
		wanted []string
	}{
		// {"()())()", []string{"(())()", "()()()"}},
		// {"(a)())()", []string{"(a())()", "(a)()()"}},
		// {")(", []string{""}},
		// {"a", []string{"a"}},
		// {"()", []string{"()"}},
		{"((((((((((((((((((((aaaaa", []string{"aaaaa"}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("s: %s", c.s), func(t *testing.T) {
			assert.ElementsMatch(t, c.wanted, removeInvalidParentheses(c.s))
		})
	}
}
