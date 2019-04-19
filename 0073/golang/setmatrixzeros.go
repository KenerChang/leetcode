package setmatrixzeros

// setZeroes iterates matrix and mark the first row and first column
// if this row or this column should be zeros
// after the matrix is marked, it then re-iterate the matrix and set zeros
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	if len(matrix[0]) == 0 {
		return
	}

	rowNum := len(matrix)
	colNum := len(matrix[0])

	firstRowZero := false
	firstColZero := false

	for rowIdx, row := range matrix {
		for colIdx, vol := range row {
			if vol == 0 {
				matrix[0][colIdx] = 0
				matrix[rowIdx][0] = 0

				if colIdx == 0 {
					firstColZero = true
				}

				if rowIdx == 0 {
					firstRowZero = true
				}
			}
		}
	}

	for i := 1; i < rowNum; i++ {
		for j := 1; j < colNum; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		// set frist to zeros
		for i := 0; i < colNum; i++ {
			matrix[0][i] = 0
		}
	}

	if firstColZero {
		for i := 0; i < rowNum; i++ {
			matrix[i][0] = 0
		}
	}

}
