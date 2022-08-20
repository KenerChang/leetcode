package largestrectangleinhistogram

func largestRectangleArea(heights []int) int {
	// use a stack to solve if problem
	// if heights[i] >= heights[stack[top]], then put it into the stack
	// else we pop stack and compute max area until heights[i] > heights[stack[top]]

	stack := []int{-1}
	result := 0
	for i, height := range heights {
		for stack[len(stack)-1] != -1 && height < heights[stack[len(stack)-1]] {
			// do pop
			topIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// compute area
			areaWidth := i - stack[len(stack)-1] - 1
			areaHeight := heights[topIdx]
			result = max(result, areaWidth*areaHeight)
		}
		stack = append(stack, i)
	}

	// if stack is not empty, pop element and compute max area
	lastPosition := len(heights)
	for len(stack) > 0 && stack[len(stack)-1] != -1 {
		topIdx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// compute area
		areaWidth := lastPosition - stack[len(stack)-1] - 1
		areaHeight := heights[topIdx]
		result = max(result, areaWidth*areaHeight)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
