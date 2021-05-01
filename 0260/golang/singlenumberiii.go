package singlenumberiii

func singleNumber(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}

	var xorResult int

	for _, num := range nums {
		xorResult = xorResult ^ num
	}

	bitInDiff := xorResult & -xorResult

	results := make([]int, 2)
	for _, num := range nums {
		if num&bitInDiff == 0 {
			results[0] ^= num
		} else {
			results[1] ^= num
		}
	}

	return results
}
