package sentencescreenfitting

func wordsTyping(sentence []string, rows int, cols int) int {
	totalGrids := rows * cols

	var count int
	var colPosition int = -1
	for totalGrids > 0 {
		for _, word := range sentence {
			if len(word) > cols {
				return 0
			}

			if len(word) > totalGrids {
				return count
			}

			if colPosition+len(word) < cols {
				colPosition += len(word)
				totalGrids -= len(word)
			} else {
				// swtich to a new row
				totalGrids -= cols - 1 - colPosition
				colPosition = len(word) - 1
				totalGrids -= len(word)
			}

			// fmt.Printf("word: %s, colPosition: %d\n", word, colPosition)

			if colPosition < cols-1 {
				// put a space
				totalGrids--
				colPosition++
			}
			// fmt.Printf("word: %s, colPosition: %d\n", word, colPosition)
		}

		// fmt.Printf("totalGrids: %d, colPosition: %d\n", totalGrids, colPosition)

		count++
	}

	return count
}
