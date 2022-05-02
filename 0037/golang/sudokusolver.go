package sudokusolver

type Position struct {
	Row        int
	Column     int
	Candidates []byte
}

func solveSudoku(board [][]byte) {
	solved := solveSudokuRecursive(board)
	if !solved {
		panic("not solved")
	}

	// fmt.Printf("board: %+v\n", board)
}

func solveSudokuRecursive(board [][]byte) bool {
	// we use buckets to store unsolved position

	// for each iteration we keep possible digits of each row and each column
	// we also count possible candidates of each unsolved position

	// after each iteration, we fill positions which have only 1 candiate
	// if there have no position having only 1 cadiate, we choose one position which
	// has mupltiple candiate and try if it works.

	changes := []*Position{}

	for rowIdx, row := range board {
		rowBase := (rowIdx / 3) * 3
		for colIdx, val := range row {
			if val != '.' {
				continue
			}
			colBase := (colIdx / 3) * 3

			// get available row candidates
			rowCandidates := getRowCandiates(board, rowIdx)

			// get available column candidates
			colCandidates := getColumnCandidates(board, colIdx)

			// get available sub squre candidates
			subSCandidates := subSqureCandidates(board, rowBase, colBase)

			// get candidates of this position
			candidates := intersection(rowCandidates, colCandidates)
			candidates = intersection(candidates, subSCandidates)

			if len(candidates) == 0 {
				rollback(board, changes)
				return false
			} else if len(candidates) == 1 {
				// only one candidate, fill position
				board[rowIdx][colIdx] = candidates[0]

				changes = append(changes, &Position{
					Row:    rowIdx,
					Column: colIdx,
				})
			} else {
				for _, candidate := range candidates {
					board[rowIdx][colIdx] = candidate
					if solveSudokuRecursive(board) {
						return true
					}

					board[rowIdx][colIdx] = '.'
				}
				rollback(board, changes)
				return false
			}
		}
	}
	return true
}

func subSqureCandidates(board [][]byte, rowBase, colBase int) []byte {
	candidates := map[byte]bool{
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			digit := board[rowBase+i][colBase+j]
			if digit == '.' {
				continue
			}

			delete(candidates, digit)
		}
	}

	results := []byte{}
	for digit := range candidates {
		results = append(results, digit)
	}
	return results
}

func getRowCandiates(board [][]byte, rowIdx int) []byte {
	candidates := map[byte]bool{
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	row := board[rowIdx]

	for _, col := range row {
		if col == '.' {
			continue
		}

		delete(candidates, col)
	}

	results := []byte{}
	for candidate := range candidates {
		results = append(results, candidate)
	}

	return results
}

func getColumnCandidates(board [][]byte, colIdx int) []byte {
	candidates := map[byte]bool{
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	for _, row := range board {
		if row[colIdx] == '.' {
			continue
		}
		delete(candidates, row[colIdx])
	}

	results := []byte{}
	for candidate := range candidates {
		results = append(results, candidate)
	}

	return results
}

func intersection(s1, s2 []byte) []byte {
	s1Chars := map[byte]bool{}
	for _, char := range s1 {
		s1Chars[char] = true
	}

	results := []byte{}
	for _, char := range s2 {
		if _, ok := s1Chars[char]; ok {
			results = append(results, char)
		}
	}

	return results
}

func rollback(board [][]byte, changes []*Position) {
	// fmt.Printf("changes: %+v\n", changes)
	for _, change := range changes {
		board[change.Row][change.Column] = '.'
	}
}
