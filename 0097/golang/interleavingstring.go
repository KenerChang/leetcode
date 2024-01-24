package interleavingstring

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleaveDP(s1, s2, s3, map[string]bool{})
}

func isInterleaveDP(s1 string, s2 string, s3 string, cache map[string]bool) bool {
	// If we represent all the possible combinations as a tree, then
	// our goal is to find if there is a path to leaf that both
	// s1 and s2 are consumed. In this case this problem is a graph problem
	// then we can use DFS to solve.

	// matched
	if s1 == "" && s2 == "" && s3 == "" {
		return true
	}

	// length not match
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// fmt.Printf("s1: %s, s2: %s, s3: %s\n", s1, s2, s3)

	key := fmt.Sprintf("%s,%s", s1, s2)
	if cached, ok := cache[key]; ok {
		return cached
	}

	// start with s1 substrings
	index := 0

	for index < len(s1) && s1[index] == s3[index] {
		if isInterleaveDP(getSubstring(s1, index), s2, getSubstring(s3, index), cache) {
			cache[key] = true
			return true
		}
		index++
	}

	// start with s2 substrings
	index = 0
	for index < len(s2) && s2[index] == s3[index] {
		if isInterleaveDP(s1, getSubstring(s2, index), getSubstring(s3, index), cache) {
			cache[key] = true
			return true
		}
		index++
	}

	cache[key] = false
	return false
}

func getSubstring(s string, index int) string {
	if index > len(s)-1 {
		return ""
	}

	return s[index+1:]
}
