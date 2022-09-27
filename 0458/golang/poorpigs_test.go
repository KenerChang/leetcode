package poorpigs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoorPigs(t *testing.T) {
	cases := []struct {
		name          string
		wanted        int
		buckets       int
		minutesToDie  int
		minutesToTest int
	}{
		{"test1", 5, 1000, 15, 60},
		{"test2", 2, 4, 15, 15},
		{"test3", 2, 4, 15, 30},
		{"test4", 3, 125, 1, 4},
		{"test5", 2, 8, 15, 40},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, poorPigs(c.buckets, c.minutesToDie, c.minutesToTest))
		})
	}
}
