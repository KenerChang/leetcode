package generateparentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	cases := []struct {
		target []string
		n      int
	}{
		{
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"}, 3,
		},
		{
			[]string{"()"}, 1,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.n), func(t *testing.T) {
			assert.ElementsMatch(t, c.target, generateParenthesis(c.n))
		})
	}
}
