package validnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	cases := []struct {
		s      string
		wanted bool
	}{
		{
			s:      "2",
			wanted: true,
		},
		{
			s:      "0089",
			wanted: true,
		},
		{
			s:      "-0.1",
			wanted: true,
		},
		{
			s:      "+3.14",
			wanted: true,
		},
		{
			s:      "4.",
			wanted: true,
		},
		{
			s:      "-.9",
			wanted: true,
		},
		{
			s:      "2e10",
			wanted: true,
		},
		{
			s:      "90E3",
			wanted: true,
		},
		{
			s:      "3e+7",
			wanted: true,
		},
		{
			s:      "+6e-1",
			wanted: true,
		},
		{
			s:      "53.5e93",
			wanted: true,
		},
		{
			s:      "-123.456e789",
			wanted: true,
		},
		{
			s:      "abc",
			wanted: false,
		},
		{
			s:      "1a",
			wanted: false,
		},
		{
			s:      "1e",
			wanted: false,
		},
		{
			s:      "99e2.5",
			wanted: false,
		},
		{
			s:      "--6",
			wanted: false,
		},
		{
			s:      "-+3",
			wanted: false,
		},
		{
			s:      "95a54e53",
			wanted: false,
		},
		{
			s:      "e3",
			wanted: false,
		},
		{
			s:      ".1",
			wanted: true,
		},
		{
			s:      ".1e+55",
			wanted: true,
		},
		{
			s:      "46.e3",
			wanted: true,
		},
		{
			s:      "46.e+3",
			wanted: true,
		},
		{
			s:      "46.e-3",
			wanted: true,
		},
	}
	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Equal(t, c.wanted, isNumber(c.s))
		})
	}
}
