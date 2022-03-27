package shortestworddistance

import "math"

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	// iterate wordsDict

	distance := math.MaxInt32
	var word1Pos, word2Pos int = -1, -1
	for i, word := range wordsDict {
		// if word == word1, check if distance gets less
		if word == word1 {
			word1Pos = i

			if word2Pos >= 0 {
				distance = min(distance, word1Pos-word2Pos)
			}
		}

		// if word == word2, check if distance gets less
		if word == word2 {
			word2Pos = i
			if word1Pos >= 0 {
				distance = min(distance, word2Pos-word1Pos)
			}
		}
	}

	return distance
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
