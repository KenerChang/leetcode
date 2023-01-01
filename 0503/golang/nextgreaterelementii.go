package nextgreaterelementii

func nextGreaterElements(nums []int) []int {
	// we could iterate the nums twice, for a number num, the numbers appear before num
	// can be viewed as apperaing after num, it equals to find next greater element of num

	stack := [][]int{}
	results := make([]int, len(nums))
	for i := range nums {
		results[i] = -1
	}

	for i := 0; i < 2*len(nums); i++ {
		pos := i % len(nums)
		num := nums[pos]

		for len(stack) > 0 && num > stack[len(stack)-1][1] {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			results[top[0]] = num
		}

		stack = append(stack, []int{pos, num})
	}
	return results
}
