package smallestrectangleenclosingblackpixels

import (
	"math"
)

func minArea(image [][]byte, x int, y int) int {
	var minX, maxX, minY, maxY int = x, x, y, y

	minX = min(minX, findMinX(image, x))
	maxX = max(maxX, findMaxX(image, x))

	minY = min(minY, findMinY(image, y))
	maxY = max(maxY, findMaxY(image, y))

	// fmt.Printf("maxX: %d, minX: %d, maxY: %d, minY: %d\n", maxX, minX, maxY, minY)

	return (maxX - minX + 1) * (maxY - minY + 1)
}

func findMinX(image [][]byte, maxX int) int {
	var l, r int = 0, maxX - 1

	var result int = math.MaxInt
	for l <= r {
		mid := (l + r) / 2

		// iterate the row to see if any black pixel
		var hasBlack bool
		for i := 0; i < len(image[0]); i++ {
			if image[mid][i] == '1' {
				hasBlack = true
				break
			}
		}

		if hasBlack {
			result = min(result, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result
}

func findMaxX(image [][]byte, minX int) int {
	var l, r int = minX + 1, len(image) - 1

	var result int = math.MinInt

	for l <= r {
		mid := (l + r) / 2

		var hasBlack bool
		for i := 0; i < len(image[0]); i++ {
			if image[mid][i] == '1' {
				hasBlack = true
				break
			}
		}

		if hasBlack {
			result = max(result, mid)
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return result
}

func findMinY(image [][]byte, maxY int) int {
	var l, r int = 0, maxY - 1
	var result = math.MaxInt

	for l <= r {
		mid := (l + r) / 2

		var hasBlack bool
		for i := 0; i < len(image); i++ {
			if image[i][mid] == '1' {
				hasBlack = true
				break
			}
		}

		if hasBlack {
			result = min(result, mid)

			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result
}

func findMaxY(image [][]byte, minY int) int {
	var l, r int = minY + 1, len(image[0]) - 1
	var result = math.MinInt

	for l <= r {
		mid := (l + r) / 2

		var hasBlack bool
		for i := 0; i < len(image); i++ {
			if image[i][mid] == '1' {
				hasBlack = true
				break
			}
		}

		if hasBlack {
			result = max(result, mid)

			l = mid + 1
		} else {
			r = mid - 1
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
