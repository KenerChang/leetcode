package wordladder

// ladderLength use BFS to search the graph which can from beginWord to endWord
// if there are no way to link to endWord
// it will return 0
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	endWordExists := false
	patterns := map[string][]string{}
	for _, word := range wordList {
		if word == endWord {
			endWordExists = true
		}

		for i := range word {
			pattern := word[:i] + "*" + word[i+1:]

			if _, found := patterns[pattern]; !found {
				patterns[pattern] = []string{}
			}

			patterns[pattern] = append(patterns[pattern], word)
		}
	}

	if !endWordExists {
		return 0
	}

	type step struct {
		word  string
		steps int
	}

	queue := []step{
		step{
			word:  beginWord,
			steps: 1,
		},
	}

	for len(queue) > 0 {
		currentStep := queue[0]
		word := currentStep.word
		queue = queue[1:]

		// find next step candidates
		for i := range word {
			pattern := word[:i] + "*" + word[i+1:]

			if candidates, found := patterns[pattern]; found {
				for _, canddiate := range candidates {
					if canddiate == endWord {
						return currentStep.steps + 1
					}

					queue = append(queue, step{
						word:  canddiate,
						steps: currentStep.steps + 1,
					})
				}

				// delete visited pattern
				delete(patterns, pattern)
			}
		}
	}

	// not found
	return 0
}
