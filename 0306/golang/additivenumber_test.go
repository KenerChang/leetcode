package additivenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAdditiveNumber(t *testing.T) {
	cases := []struct {
		wanted bool
		num    string
	}{
		{true, "112358"},
		{true, "199100199"},
		{true, "123"},
		{false, "12"},
		{true, "101121"},
		{true, "011"},
		{true, "000"},
		{false, "124"},
		{true, "11235813213455"},
		{false, "1102"},
		{true, "10203050"},
		{true, "011235"},
		{true, "0112358"},
		{false, "01011"},
		{false, "001010"},
		{false, "10"},
	}

	for _, c := range cases {
		t.Run(c.num, func(t *testing.T) {
			assert.Equal(t, c.wanted, isAdditiveNumber(c.num))
		})
	}
}
