package reversewordsinastring

import (
	"strings"
)

func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	tokens := strings.Fields(s)

	// revert
	for i, j := 0, len(tokens)-1; j > i; i, j = i+1, j-1 {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	}

	return strings.Join(tokens, " ")
}
