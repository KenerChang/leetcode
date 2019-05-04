package searchinrotatedsortedarray

func search(nums []int, target int) bool {
	// at each iteration, the int slice can split into two parts
	// one part is a monotonic numbers, another will has pivot
	// so for each iteration, we can if target is in monotonic part
	// if target is not in monotonic part, we try to find pivot part and do the search
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return nums[0] == target
	}

	found := false
	left := 0
	right := len(nums) - 1
	for !found {
		middle := (left + right) / 2

		if nums[middle] == target || nums[left] == target || nums[right] == target {
			found = true
			break
		}

		if middle == left || middle == right {
			break
		}

		if nums[middle+1] == target || nums[middle-1] == target {
			found = true
			break
		}

		// left part is monotonic increasing
		if nums[left] < nums[middle-1] {
			if nums[left] <= target && nums[middle-1] >= target {
				right = middle - 1
				continue
			}
		}

		// right part is monotonic increasing
		if nums[middle+1] < nums[right] {
			if nums[middle+1] <= target && nums[right] >= target {
				left = middle + 1
				continue
			}
		}

		// left part has pivot
		if nums[left] > nums[middle-1] {
			right = middle - 1
			continue
		}

		// right part has pivot
		if nums[middle+1] > nums[right] {
			left = middle + 1
			continue
		}

		// check if left part has pivot
		if nums[middle-1] == nums[left] {
			has := hasPivot(nums[left:middle])

			if has {
				right = middle - 1
				continue
			}
		}

		if nums[middle+1] == nums[right] {
			has := hasPivot(nums[middle+1 : right+1])

			if has {
				left = middle + 1
				continue
			}
		}
		break
	}
	return found
}

func hasPivot(nums []int) bool {
	has := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			has = true
			break
		}
	}
	return has
}
