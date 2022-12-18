package mostcommonword

import "strings"

func clean(paragraph string) string {
	// clean !?',;.

	paragraph = strings.ToLower(paragraph)
	paragraph = strings.ReplaceAll(paragraph, "!", " ")
	paragraph = strings.ReplaceAll(paragraph, "?", " ")
	paragraph = strings.ReplaceAll(paragraph, "'", " ")
	paragraph = strings.ReplaceAll(paragraph, ",", " ")
	paragraph = strings.ReplaceAll(paragraph, ";", " ")
	paragraph = strings.ReplaceAll(paragraph, ".", " ")

	return paragraph
}

func mostCommonWord(paragraph string, banned []string) string {
	// build banned map
	bannedMap := map[string]bool{}
	for _, ban := range banned {
		bannedMap[ban] = true
	}

	paragraph = clean(paragraph)
	tokens := strings.Split(paragraph, " ")

	var result string
	var maxFrequency int
	tokenMap := map[string]int{}
	for _, token := range tokens {
		if token == "" {
			continue
		}

		if _, ok := bannedMap[token]; ok {
			continue
		}

		tokenMap[token]++

		if tokenMap[token] > maxFrequency {
			result = token
			maxFrequency = tokenMap[token]
		}
	}
	return result
}
