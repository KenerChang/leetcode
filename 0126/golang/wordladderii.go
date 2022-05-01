package wordladderii

type Node struct {
	Level  int
	Val    string
	Parent *Node
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// use BFS
	// each time we pop a node from stack, and put all its children (distance = 1)
	// into the stack if:
	// 1. the node has not been added, or
	// 2. current level it not larget than nodes' level it has been put to stack

	stack := []*Node{
		{
			Level:  0,
			Val:    beginWord,
			Parent: nil,
		},
	}

	// keep word level
	usedWords := map[string]int{}

	wordCandidates := map[string]bool{}
	for _, word := range wordList {
		wordCandidates[word] = true
	}

	results := [][]string{}
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		// check if endWord
		nodeWord := node.Val
		if nodeWord == endWord {
			result := []string{}

			for node != nil {
				result = append(result, node.Val)
				node = node.Parent
			}

			reverse(result)
			results = append(results, result)
			continue
		}

		for word := range wordCandidates {
			// check if this word can be used
			if !canUseWord(word, endWord, node.Level+1, usedWords) {
				delete(wordCandidates, word)
				continue
			}

			if distance(nodeWord, word) > 1 {
				continue
			}

			stack = append(stack, &Node{
				Val:    word,
				Level:  node.Level + 1,
				Parent: node,
			})

			usedWords[word] = node.Level + 1
		}
	}

	// fmt.Printf("results: %+v\n", results)
	return results
}

func canUseWord(word, endWord string, level int, usedWords map[string]int) bool {
	if prevLevel, ok := usedWords[word]; ok {
		return level <= prevLevel
	} else {
		// word has not been used
		return true
	}
}

func distance(s1, s2 string) int {
	result := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			result++
		}
	}

	return result
}

func reverse(ss []string) {
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
}
