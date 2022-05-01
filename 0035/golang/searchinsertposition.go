package searchinsertposition

func searchInsert(nums []int, target int) int {
	// do binary search to find first postion i which nums[i] >= target
	start := 0
	end := len(nums)

	for start < end {
		middle := (end + start) / 2

		if nums[middle] >= target {
			// search left part
			end = middle
		} else {
			// search right
			start = middle + 1
		}
	}

	return start
}
