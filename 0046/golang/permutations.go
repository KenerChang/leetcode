package permutations

func permute(nums []int) [][]int {
	// use recursive
	// iterate nums and choose one number to put at first
	// the remain numbers are put to permute funtion

	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{
			[]int{
				nums[0],
			},
		}
	}

	results := [][]int{}
	lastIdx := len(nums) - 1
	for idx, num := range nums {
		var remains []int
		if idx == 0 {
			remains = nums[1:]
		} else if idx == lastIdx {
			remains = nums[0:lastIdx]
		} else {
			remains = append([]int{}, nums[0:idx]...)
			remains = append(remains, nums[idx+1:]...)
		}

		subResults := permute(remains)
		for _, subResult := range subResults {
			result := append([]int{num}, subResult...)
			results = append(results, result)
		}
	}
	return results
}