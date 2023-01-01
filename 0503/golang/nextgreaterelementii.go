package nextgreaterelementii

func nextGreaterElements(nums []int) []int {
	// find the left most greater value for each num
	// find the next greter value for each num
	// if a num has next greater value => next greater value
	// else if a num has left most greater value => left most greater value
	// else -1

	leftMostGreaters := make([]int, len(nums))
	leftMostGreaters[0] = -1

	queue := []int{nums[0]} // montonic ascending queue
	for i, num := range nums {
		if num > queue[len(queue)-1] {
			queue = append(queue, num)
			leftMostGreaters[i] = -1
		} else if num < queue[0] {
			leftMostGreaters[i] = queue[0]
		} else {
			var found bool
			for j := 0; j < len(queue); j++ {
				if queue[j] > num {
					leftMostGreaters[i] = queue[j]
					found = true
					break
				}

				if !found {
					leftMostGreaters[i] = -1
				}
			}
		}
	}

	// fmt.Printf("leftMostGreaters: %v\n", leftMostGreaters)

	stack := [][]int{} // 0: num index 1: val // monotonic ascending
	results := make([]int, len(nums))
	for i, num := range nums {
		for len(stack) > 0 && num > stack[len(stack)-1][1] {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			results[top[0]] = num
		}

		stack = append(stack, []int{i, num})
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		results[top[0]] = leftMostGreaters[top[0]]
	}

	return results
}
