package searcha2dmatrixii

func searchMatrix(matrix [][]int, target int) bool {
	mLen := len(matrix)
	if mLen == 0 {
		return false
	}

	nLen := len(matrix[0])
	if nLen == 0 {
		return false
	}

	return searchMatrixRecursive(matrix, 0, 0, mLen-1, nLen-1, target)
}

func searchMatrixRecursive(matrix [][]int, startX, startY, endX, endY, target int) bool {
	// fmt.Printf("startX: %d, startY: %d, endX: %d, endY: %d\n", startX, startY, endX, endY)
	if startX > endX || startY > endY {
		return false
	}

	middleX := (startX + endX) / 2
	middleY := (startY + endY) / 2
	middle := matrix[middleX][middleY]

	if middle == target {
		return true
	}

	if middle > target {
		return searchMatrixRecursive(matrix, startX, startY, endX, middleY-1, target) || searchMatrixRecursive(matrix, startX, middleY, middleX-1, endY, target)
	}
	return searchMatrixRecursive(matrix, middleX+1, startY, endX, endY, target) || searchMatrixRecursive(matrix, startX, middleY+1, middleX, endY, target)
}
