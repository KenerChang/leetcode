package missingranges

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	if len(nums) == 0 {
		return [][]int{{lower, upper}}
	}

	results := [][]int{}
	if lower < nums[0] {
		results = append(results, []int{lower, nums[0] - 1})
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			results = append(results, []int{nums[i] + 1, nums[i+1] - 1})
		}
	}

	if upper > nums[len(nums)-1] {
		results = append(results, []int{nums[len(nums)-1] + 1, upper})
	}
	return results
}
