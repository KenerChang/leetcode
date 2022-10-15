package fizzbuzz

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		wanted []string
		n      int
	}{
		{[]string{"1", "2", "Fizz"}, 3},
		{[]string{"1", "2", "Fizz", "4", "Buzz"}, 5},
		{[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, 15},
	}

	for _, c := range cases {
		t.Run(strconv.Itoa(c.n), func(t *testing.T) {
			assert.Equal(t, c.wanted, fizzBuzz(c.n))
		})
	}
}
