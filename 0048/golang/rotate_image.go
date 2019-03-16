package rotate

import "math"

func rotate(matrix [][]int)  {
	// rotate 90 degree needs two steps
	// 1. transpose matrix
	// 2. reverse each row
	size := len(matrix)
	lastIdx := size - 1

	// the row i of new matrix is reverse of column i of original matrix 
	var swap int
	for i, row := range matrix {
		for j := i+1; j < size; j++ {
			swap = row[j]
			matrix[i][j], matrix[j][i] = matrix[j][i], swap
		}
	}

	// reverse each row
	middle := int(math.Ceil(float64(size)/2))
	for i, row := range matrix {
		for j := 0; j < middle; j++ {
			swap = row[j]
			matrix[i][j], matrix[i][lastIdx-j] = matrix[i][lastIdx-j], swap 
		}
	}
}