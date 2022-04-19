package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	results := []int{}
	// build dictionary
	wordDict := map[string]int{}
	for _, word := range words {
		if _, ok := wordDict[word]; !ok {
			wordDict[word] = 0
		}

		wordDict[word]++
	}

	wordLength := len(words[0])
	substringLength := wordLength * len(words)

	// iterate s
	for i := 0; i+substringLength <= len(s); i++ {
		usedWords := map[string]int{}
		// fmt.Printf("i: %d\n", i)

		word := s[i : i+wordLength]
		// if word can not be used, continue to next char
		if !canUse(word, wordDict, usedWords) {
			continue
		}

		// fmt.Printf("word: %s\n", word)

		usedWords[word]++

		// check if a substring
		substring := word
		for j := i + wordLength; j < i+substringLength; j += wordLength {
			// fmt.Printf("i: %d, j: %d, wordLength: %d\n", i, j, wordLength)
			nextWord := s[j : j+wordLength]
			// fmt.Printf("nextWord: %s\n", nextWord)
			if !canUse(nextWord, wordDict, usedWords) {
				break
			}

			usedWords[nextWord]++
			substring += word
		}

		if len(substring) == substringLength {
			results = append(results, i)
		}
	}
	return results
}

func canUse(word string, wordDict, usedWords map[string]int) bool {
	total, inDict := wordDict[word]
	if !inDict {
		return false
	}

	usedTimes, used := usedWords[word]
	return !used || usedTimes < total
}
