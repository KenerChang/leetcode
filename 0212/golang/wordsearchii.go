package wordsearchii

import (
// "fmt"
// "math"
)

const a = 'a'

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	charPositions := make([][]int, 26)
	for idx, word := range words {
		if word == "" {
			continue
		}

		charIdx := word[0] - a
		charPositions[charIdx] = append(charPositions[charIdx], idx)
	}

	results := []string{}
	for rIdx, row := range board {
		for cIdx, char := range row {
			charIdx := char - a
			if len(charPositions[charIdx]) == 0 {
				continue
			}

			for idx, wordIdx := range charPositions[charIdx] {
				word := words[wordIdx]
				if findWordsRecursive(board, word, rIdx, cIdx) {
					results = append(results, word)
					charPositions[charIdx] = append(charPositions[charIdx][:idx], charPositions[charIdx][idx+1:]...)
					continue
				}
			}
		}
	}
	return results
}

func findWordsRecursive(board [][]byte, word string, rIdx, cIdx int) bool {
	// fmt.Printf("word: %s, searchFrom: %d\n", word, searchFrom)
	// fmt.Printf("visited: %+v\n", visited)
	if word == "" {
		return true
	}

	if rIdx < 0 || cIdx < 0 || rIdx >= len(board) || cIdx >= len(board[0]) {
		return false
	}

	if board[rIdx][cIdx] == '#' || word[0] != board[rIdx][cIdx] {
		return false
	}

	c := board[rIdx][cIdx]
	board[rIdx][cIdx] = '#'
	defer func() {
		board[rIdx][cIdx] = c
	}()

	if findWordsRecursive(board, word[1:], rIdx+1, cIdx) {
		return true
	}

	if findWordsRecursive(board, word[1:], rIdx-1, cIdx) {
		return true
	}

	if findWordsRecursive(board, word[1:], rIdx, cIdx+1) {
		return true
	}

	if findWordsRecursive(board, word[1:], rIdx, cIdx-1) {
		return true
	}

	return false
}
