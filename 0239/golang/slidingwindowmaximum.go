package slidingwindowmaximum

func maxSlidingWindow(nums []int, k int) []int {
	// We can use monotonic queue to solve this problem.
	// We use decreasing queue to store max value in sliding windon on the first position.
	// Every time we move the window, we pop out elements which are less than the coming num
	// and then put the coming num into the queue.

	if len(nums) <= k {
		maxNum, _ := max(nums...)
		return []int{maxNum}
	}

	// init
	queue := [][]int{}
	for i := 0; i < k; i++ {
		num := nums[i]

		// pop all elements which are less than num
		for len(queue) > 0 && (queue[len(queue)-1][0] < num) {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, []int{num, i})
	}

	results := []int{queue[0][0]}
	for i := 1; i+k <= len(nums); i++ {
		num := nums[i+k-1]
		// pop all elements which are less than num and elemnts which are out of the window
		for len(queue) > 0 && (queue[len(queue)-1][0] < num) {
			queue = queue[:len(queue)-1]
		}

		for len(queue) > 0 && (queue[0][1] < i) {
			queue = queue[1:]
		}

		queue = append(queue, []int{num, i + k - 1})
		results = append(results, queue[0][0])
	}

	return results
}

func max(nums ...int) (int, int) {
	result := nums[0]
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
			idx = i
		}
	}

	return result, idx
}
