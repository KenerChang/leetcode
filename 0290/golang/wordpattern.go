package wordpattern

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	// split s by space
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	patternToWordMap := map[byte]string{}
	wordToPatternMap := map[string]byte{}

	for i, word := range words {
		wordPattern, ok := wordToPatternMap[word]
		if ok && wordPattern != pattern[i] {
			return false
		}

		patternWord, ok := patternToWordMap[pattern[i]]
		if ok && patternWord != word {
			return false
		}

		wordToPatternMap[word] = pattern[i]
		patternToWordMap[pattern[i]] = word
	}

	return true
}
