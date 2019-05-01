package wordsearch

// exist searches if a word is sequence character in board
func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}

	wordBytes := []byte(word)
	firstChar := wordBytes[0]

	result := false
	for i, row := range board {
		if result {
			break
		}

		for j, char := range row {
			if result {
				break
			}

			if char == firstChar {
				result = startWith(&board, i, j, wordBytes)
				if result {
					break
				}
			}
		}
	}
	return result
}

// startWith check sequence exists by given start point
func startWith(board *[][]byte, x, y int, target []byte) bool {
	if x < 0 || x >= len(*board) {
		return false
	}

	if y < 0 || y >= len((*board)[0]) {
		return false
	}

	if (*board)[x][y] == '-' {
		return false
	}

	if (*board)[x][y] != target[0] {
		return false
	}

	if len(target) == 1 {
		return true
	}

	c := (*board)[x][y]
	(*board)[x][y] = '-'

	// check up, down, left, right
	result := startWith(board, x-1, y, target[1:len(target)]) || startWith(board, x+1, y, target[1:len(target)]) || startWith(board, x, y-1, target[1:len(target)]) || startWith(board, x, y+1, target[1:len(target)])

	(*board)[x][y] = c

	return result
}
