package subdomainvisitcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubdomainVisits(t *testing.T) {
	cases := []struct {
		name      string
		wanted    []string
		cpdomains []string
	}{
		{"case 1", []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"}, []string{"9001 discuss.leetcode.com"}},
		{"case 2", []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}, []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}},
		{"case 2", []string{"15 org", "5 wiki.org"}, []string{"10 org", "5 wiki.org"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.wanted, subdomainVisits(c.cpdomains))
		})
	}
}
