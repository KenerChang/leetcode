package generalizedabbreviation

import (
	"strconv"
)

func backtrack(word, prefix string, start, k int, results []string) []string {
	// fmt.Printf("word: %s, prefix: %s, start: %d\n", word, prefix, start)
	if start >= len(word) {
		if k > 0 {
			prefix += strconv.Itoa(k)
		}

		results = append(results, prefix)
		return results
	}

	// the branch that word[start] is abbreviated
	results = backtrack(word, prefix, start+1, k+1, results)

	// the branch that word[start] is kept
	var newPrefix string = prefix
	if k > 0 {
		newPrefix += strconv.Itoa(k)
	}
	results = backtrack(word, newPrefix+string(word[start]), start+1, 0, results)
	return results
}

func generateAbbreviations(word string) []string {
	// Constraints:
	// - the abbreviations can not overlap
	// - the abbreviations can not adjacent

	return backtrack(word, "", 0, 0, []string{})
}
