package shortestpalindrome

import (
// "fmt"
)

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	i, j := 0, len(s)-1

	// find max common match
	for j >= 0 {
		if s[i] == s[j] {
			i++
		}
		j--
	}

	if i == len(s) {
		return s
	}

	// shortest palidrome is composed of reverse of non match part + shortestPalindrome of match part + non match part
	return reverse(s[i:]) + shortestPalindrome(s[:i]) + s[i:]
}

func reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
