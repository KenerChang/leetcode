package reversewordsinastring

import (
	"strings"
)

func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	tokens := strings.Split(s, " ")

	results := []string{}
	for _, token := range tokens {
		if token != "" {
			results = append(results, token)
		}
	}

	result := ""
	for i := len(results) - 1; i >= 0; i-- {
		result += results[i] + " "
	}

	if len(result) > 0 {
		result = result[0 : len(result)-1]
	}
	return result
}
