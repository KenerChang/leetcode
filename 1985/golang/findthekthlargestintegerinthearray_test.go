package findthekthlargestintegerinthearray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargestNumber(t *testing.T) {
	cases := []struct {
		name   string
		wanted string
		given  []string
		k      int
	}{
		{"case 1", "3", []string{"3", "6", "7", "10"}, 4},
		{"case 2", "2", []string{"2", "21", "12", "1"}, 3},
		{"case 3", "0", []string{"0", "0"}, 2},
		{"case 4", "3067240", []string{"83348", "44437", "858441", "9665240415", "57338484", "299", "4083", "1165", "22082677", "5540", "48", "3", "3067240", "5", "2547887287"}, 5},
		{"case 5", "8", []string{"61718", "1262", "8", "696", "6", "75780"}, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, kthLargestNumber(c.given, c.k))
		})
	}
}
