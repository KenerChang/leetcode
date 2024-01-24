package interleavingstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	cases := []struct {
		s1     string
		s2     string
		s3     string
		wanted bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"aabcc", "", "aabcc", true},
		{"", "dbbca", "dbbca", true},
		{"aabcc", "dbbca", "aabccdbbca", true},
		{"aabcc", "dbbca", "dbbcaaabcc", true},
		{"aaa", "bbb", "ababab", true},
		{"aaa", "bbb", "bababa", true},
		{"aaa", "bbb", "aaabbbb", false},
		{"bcbccabcccbcbbbcbbacaaccccacbaccabaccbabccbabcaabbbccbbbaa", "ccbccaaccabacaabccaaccbabcbbaacacaccaacbacbbccccbac", "bccbcccabbccaccaccacbacbacbabbcbccbaaccbbaacbcbaacbacbaccaaccabcaccacaacbacbacccbbabcccccbababcaabcbbcccbbbaa", false},
		{"aabd", "abdc", "aabdabcd", true},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("s1: %s, s2: %s, s3: %s", c.s1, c.s2, c.s3), func(t *testing.T) {
			assert.Equal(t, c.wanted, isInterleave(c.s1, c.s2, c.s3))
		})
	}
}
