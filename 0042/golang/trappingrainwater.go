package trappingrainwater

func trap(heights []int) int {
	// use two pointers to keep left position and right position
	// use leftMax and rightMax to keep the max height of left and right
	var left int = 0
	var right int = len(heights) - 1
	var leftMax, rightMax int = 0, 0

	var result int = 0
	for left < right {
		// if heights[left] < heights[right]
		if heights[left] < heights[right] {
			// if heights[left] < leftMax, add leftMax - heights[left] to result
			if heights[left] < leftMax {
				result += leftMax - heights[left]
			} else {
				leftMax = heights[left]
			}
			// left++
			left++
		} else {
			// if heights[right] < rightMax, add rightMax - heights[right] to result
			if heights[right] < rightMax {
				result += rightMax - heights[right]
			} else {
				rightMax = heights[right]
			}
			// right--
			right--
		}
	}

	return result
}
