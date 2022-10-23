package analyzeuserwebsitevisitpattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostVisitedPattern(t *testing.T) {
	cases := []struct {
		name       string
		wanted     []string
		usernames  []string
		timestamps []int
		websites   []string
	}{
		// {
		// 	"case 1",
		// 	[]string{"home", "about", "career"},
		// 	[]string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"},
		// 	[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// 	[]string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"},
		// },
		{
			"case 2",
			[]string{"a", "b", "a"},
			[]string{"ua", "ua", "ua", "ub", "ub", "ub"},
			[]int{1, 2, 3, 4, 5, 6},
			[]string{"a", "b", "a", "a", "b", "c"},
		},
		// {
		// 	"case 3",
		// 	[]string{"a", "b", "a"},
		// 	[]string{"ua", "ua", "ua"},
		// 	[]int{1, 3, 2},
		// 	[]string{"a", "a", "b"},
		// },
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.wanted, mostVisitedPattern(c.usernames, c.timestamps, c.websites))
		})
	}
}
