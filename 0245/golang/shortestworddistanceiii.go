package shortestworddistanceiii

import "math"

func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	if word1 != word2 {
		return shortestDistanceDifferertWords(wordsDict, word1, word2)
	}
	return shortestDistanceSameWord(wordsDict, word1)
}

func shortestDistanceDifferertWords(wordsDict []string, word1 string, word2 string) int {
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

func shortestDistanceSameWord(wordsDict []string, target string) int {
	// iterate wordsDict

	distance := math.MaxInt32
	var maxPos, minPos int = -1, -1
	for i, word := range wordsDict {
		// if word != target, continue
		if word != target {
			continue
		}

		// if target occurs at first time, set to minPos, then conintue, no need to count distance
		if minPos < 0 {
			minPos = i
			continue
		} else if maxPos < 0 {
			// if target occurs at second time, set to maxPos, also set distance
			maxPos = i
			distance = maxPos - minPos
		} else {
			// target occurs more than twice, check if distance changes, and update maxPos, minPos

			distance = min(distance, i-maxPos)

			minPos = maxPos
			maxPos = i
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
