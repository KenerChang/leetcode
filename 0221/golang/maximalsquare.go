package maximalsquare

func maximalSquare(matrix [][]byte) int {
	// one := '1'
	maxSquare := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {

				if i+maxSquare-1 >= len(matrix) || j+maxSquare-1 >= len(matrix[0]) {
					continue
				}

				size := 1
				for {
					result := isSquare(matrix, i, j, size+1)
					if !result {
						if size > maxSquare {
							maxSquare = size
						}
						break
					}
					size++
				}
			}
		}
	}

	return maxSquare * maxSquare
}

func isSquare(matrix [][]byte, x, y, size int) bool {
	if x+(size-1) >= len(matrix) {
		return false
	}

	if y+(size-1) >= len(matrix[0]) {
		return false
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[x+i][y+j] == '0' {
				return false
			}
		}
	}
	return true
}
