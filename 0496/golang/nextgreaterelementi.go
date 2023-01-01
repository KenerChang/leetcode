package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums2Map := map[int]int{}

	// build map
	stack := []int{}
	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			nums2Map[top] = num
		}

		stack = append(stack, num)
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		nums2Map[top] = -1
	}

	// build result
	results := make([]int, len(nums1))
	for i := range nums1 {
		results[i] = nums2Map[nums1[i]]
	}

	return results
}
