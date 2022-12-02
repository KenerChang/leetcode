package minimumremovetomakevalidparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	cases := []struct {
		wanted []string
		s      string
	}{
		{[]string{"lee(t(c)o)de", "lee(t(co)de)", "lee(t(c)ode)"}, "lee(t(c)o)de"},
		{[]string{"ab(c)d"}, "a)b(c)d"},
		{[]string{""}, "))(("},
		{[]string{"()"}, "()"},
		{[]string{"(abcd)"}, "(abcd)"},
		{[]string{"abcd"}, "(abcd"},
		{[]string{"abcd"}, "abcd)"},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Contains(t, c.wanted, minRemoveToMakeValid(c.s))
		})
	}
}
