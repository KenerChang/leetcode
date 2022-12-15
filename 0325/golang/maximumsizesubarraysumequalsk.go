package maximumsizesubarraysumequalsk

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArrayLen(nums []int, k int) int {
	// If we do prefix sum first, we can reduce this problem to two sum
	prefixsums := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefixsums[i+1] = prefixsums[i] + nums[i]
	}

	var result int
	numMap := map[int]int{}
	for i, num := range prefixsums {
		remain := num - k
		if l, ok := numMap[remain]; ok {
			result = max(result, i-l)
		}

		if _, ok := numMap[num]; !ok {
			numMap[num] = i
		}
	}

	return result
}
