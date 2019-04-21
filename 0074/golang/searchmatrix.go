package searchmatrix

// searchMatrix searches if target is int the matrix
// it searches a target by two steps
// first, it does binary search to decide that target should in which row
// it then does binary search again to see if target is in the row
func searchMatrix(matrix [][]int, target int) bool {
	rowNum := len(matrix)
	if rowNum == 0 {
		return false
	}

	colNum := len(matrix[0])
	if colNum == 0 {
		return false
	}

	exist := false

	// decide target should in which row
	lastIdx := rowNum
	firstIdx := 0
	var middleIdx int
	var middleVal int
	for true {
		middleIdx = (lastIdx + firstIdx) / 2
		middleVal = matrix[middleIdx][0]

		if middleVal == target {
			return true
		}

		if middleIdx == lastIdx || middleIdx == firstIdx {
			break
		}

		if middleVal > target {
			lastIdx = middleIdx
		} else {
			firstIdx = middleIdx
		}
	}

	// the target should be in the row of firstIdx
	rowIdx := firstIdx
	firstIdx = 0
	lastIdx = colNum
	for true {
		middleIdx = (lastIdx + firstIdx) / 2
		middleVal = matrix[rowIdx][middleIdx]

		if middleVal == target {
			// found it
			return true
		}

		if middleIdx == lastIdx || middleIdx == firstIdx {
			break
		}

		if middleVal > target {
			lastIdx = middleIdx
		} else {
			firstIdx = middleIdx
		}
	}

	return exist
}
