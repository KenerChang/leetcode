package longeststringchain

import (
	"fmt"
	"sort"
)

type byLength []string

func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func (b byLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func longestStrChain(words []string) int {
	// To solve this problem, we first need to sort words by its length
	// After words are sorted,we can use DP to solve this problem

	// sort words
	sort.Sort(byLength(words))

	// DP
	// transform formula function
	// dp[word] = max(predecessors(word)) +1
	dp := map[string]int{}
	var result int
	for _, word := range words {
		predecessors := predecessorOf(word)

		for _, predecessor := range predecessors {
			dp[word] = max(dp[word], dp[predecessor]+1)
		}

		result = max(result, dp[word])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func predecessorOf(word string) []string {
	predecessors := make([]string, len(word))

	for i := range word {
		predecessors[i] = fmt.Sprintf("%s%s", string(word[0:i]), string(word[i+1:]))
	}

	return predecessors
}
