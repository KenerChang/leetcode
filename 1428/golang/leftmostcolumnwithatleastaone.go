package leftmostcolumnwithatleastaone

type BinaryMatrix struct {
	Get        func(int, int) int
	Dimensions func() []int
}

func GetFactory(matrix [][]int) func(int, int) int {
	return func(row int, col int) int {
		return matrix[row][col]
	}
}

func GetDimensionsFactory(matrix [][]int) func() []int {
	return func() []int {
		return []int{len(matrix), len(matrix[0])}
	}
}

func searchFirstOneInRow(matrix BinaryMatrix, rowIdx, colSize int) int {
	l, r := 0, colSize-1

	var result int = -1
	for l <= r {
		mid := (l + r) / 2
		// fmt.Printf("rowIdx: %d, l: %d, r: %d, mid: %d, val: %d, result: %d\n", rowIdx, l, r, mid, matrix.Get(rowIdx, mid), result)
		if matrix.Get(rowIdx, mid) == 1 {
			// search left
			result = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	// we search first one in each row
	// and return the smallest index

	dimensions := binaryMatrix.Dimensions()
	smallest := -1
	for i := 0; i < dimensions[0]; i++ {
		colIdx := searchFirstOneInRow(binaryMatrix, i, dimensions[1])
		if colIdx < dimensions[1] {
			// find one
			if colIdx >= 0 {
				if smallest == -1 {
					smallest = colIdx
				} else {
					smallest = min(smallest, colIdx)
				}
			}
		}
	}

	return smallest
}
