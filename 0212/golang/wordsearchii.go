package wordsearchii

import (
// "fmt"
// "math"
)

const a = 'a'

type Trie struct {
	word     string
	children [26]*Trie
}

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	root := &Trie{}
	for _, word := range words {
		node := root
		for _, char := range word {
			charIdx := char - a
			if node.children[charIdx] == nil {
				node.children[charIdx] = &Trie{}
			}

			node = node.children[charIdx]
		}
		node.word = word
	}

	results := []string{}
	for rIdx := 0; rIdx < len(board); rIdx++ {
		for cIdx := 0; cIdx < len(board[rIdx]); cIdx++ {
			findWordsRecursive(board, root, rIdx, cIdx, &results)
		}
	}
	return results
}

func findWordsRecursive(board [][]byte, root *Trie, rIdx, cIdx int, results *[]string) {
	// fmt.Printf("word: %s, searchFrom: %d\n", word, searchFrom)
	// fmt.Printf("visited: %+v\n", visited)

	if rIdx < 0 || cIdx < 0 || rIdx >= len(board) || cIdx >= len(board[0]) {
		return
	}

	c := board[rIdx][cIdx]
	charIdx := c - a
	if board[rIdx][cIdx] == '#' || root.children[charIdx] == nil {
		return
	}
	node := root.children[charIdx]

	if node.word != "" {
		// reach bottom
		*results = append(*results, node.word)
		node.word = ""
	}

	board[rIdx][cIdx] = '#'
	findWordsRecursive(board, node, rIdx+1, cIdx, results)
	findWordsRecursive(board, node, rIdx-1, cIdx, results)
	findWordsRecursive(board, node, rIdx, cIdx+1, results)
	findWordsRecursive(board, node, rIdx, cIdx-1, results)
	board[rIdx][cIdx] = c

	return
}
