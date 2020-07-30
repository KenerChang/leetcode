package wordsearchii

import (
	// "fmt"
	"math"
)

const a = 'a'

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	charPositions := make([][]int64, 26)
	for rIdx, row := range board {
		for cIdx, char := range row {
			charIdx := char - a
			pos := encodePos(rIdx, cIdx)
			charPositions[charIdx] = append(charPositions[charIdx], pos)
		}
	}

	results := []string{}
	for _, word := range words {
		if word == "" {
			continue
		}

		charIdx := word[0] - a
		if len(charPositions[charIdx]) == 0 {
			continue
		}

		for _, pos := range charPositions[charIdx] {
			if findWordsRecursive(board, word, pos, map[int64]bool{}) {
				results = append(results, word)
				break
			}
		}
	}
	return results
}

func findWordsRecursive(board [][]byte, word string, searchFrom int64, visited map[int64]bool) bool {
	// fmt.Printf("word: %s, searchFrom: %d\n", word, searchFrom)
	// fmt.Printf("visited: %+v\n", visited)

	if _, found := visited[searchFrom]; found {
		return false
	}

	if len(word) == 1 {
		return true
	}

	visited[searchFrom] = true

	row, col := decodePos(searchFrom)
	maxRow := len(board) - 1
	maxCol := len(board[0]) - 1

	if row < maxRow && board[row+1][col] == word[1] {
		pos := encodePos(row+1, col)
		if findWordsRecursive(board, word[1:], pos, visited) {
			return true
		}
	}

	if row > 0 && board[row-1][col] == word[1] {
		pos := encodePos(row-1, col)
		if findWordsRecursive(board, word[1:], pos, visited) {
			return true
		}
	}

	if col < maxCol && board[row][col+1] == word[1] {
		pos := encodePos(row, col+1)
		if findWordsRecursive(board, word[1:], pos, visited) {
			return true
		}
	}

	if col > 0 && board[row][col-1] == word[1] {
		pos := encodePos(row, col-1)
		if findWordsRecursive(board, word[1:], pos, visited) {
			return true
		}
	}

	delete(visited, searchFrom)
	return false
}

func encodePos(row, col int) int64 {
	result := int64(row)
	result = result << 32
	result += int64(col)
	return result
}

func decodePos(pos int64) (row, col int) {
	col = int(pos & math.MaxInt32)
	row = int(pos >> 32)
	return
}
