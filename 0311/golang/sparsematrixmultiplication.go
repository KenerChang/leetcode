package sparsematrixmultiplication

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	// init result matrix
	resultMatrix := make([][]int, len(mat1))
	for i := 0; i < len(resultMatrix); i++ {
		resultMatrix[i] = make([]int, len(mat2[0]))
	}

	// iterate mat1 and increase resultMatrix coresponding cells if cell value is not 0
	for rowIdx, row := range mat1 {
		for _, val := range row {
			if val != 0 {
				for colIdx := 0; colIdx < len(resultMatrix[rowIdx]); colIdx++ {
					resultMatrix[rowIdx][colIdx] = resultMatrix[rowIdx][colIdx] | 1
				}
			}
		}
	}

	// iterate mat2 and increase resultMatrix coresponding cells if cell value is not 0
	for _, row := range mat2 {
		for colIdx, val := range row {
			if val != 0 {
				for rowIdx := 0; rowIdx < len(resultMatrix); rowIdx++ {
					resultMatrix[rowIdx][colIdx] = resultMatrix[rowIdx][colIdx] | 2
				}
			}
		}
	}

	// iterate resultMatrix and compute cell values if both be increased by mat1, mat2 iteration
	for rowIdx, row := range resultMatrix {
		for colIdx, val := range row {
			if val == 3 {
				resultMatrix[rowIdx][colIdx] = getCellValue(mat1, mat2, rowIdx, colIdx)
			} else {
				// reset cell value
				resultMatrix[rowIdx][colIdx] = 0
			}
		}
	}

	return resultMatrix
}

func getCellValue(mat1, mat2 [][]int, rowIdx, colIdx int) int {
	row := mat1[rowIdx]

	var result int
	for i := 0; i < len(row); i++ {
		result += row[i] * mat2[i][colIdx]
	}

	return result
}
