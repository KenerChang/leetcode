package mostcommonword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostCommonWord(t *testing.T) {
	cases := []struct {
		wanted    string
		paragraph string
		banned    []string
	}{
		{"ball", "Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}},
		{"a", "a.", []string{""}},
		{"a", "a, a, a, a, b,b,b,c, c", []string{""}},
		{"ball", "Bob. hIt, baLl", []string{"bob", "hit"}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("paragraph:%s,banned:%v", c.paragraph, c.banned), func(t *testing.T) {
			assert.Equal(t, c.wanted, mostCommonWord(c.paragraph, c.banned))
		})
	}
}
