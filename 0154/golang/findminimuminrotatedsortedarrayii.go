package findminimuminrotatedsortedarrayii

func findMin(nums []int) int {
	return findMinRecursive(nums, 0, len(nums)-1)
}

func findMinRecursive(nums []int, start, end int) int {
	// do binary search for find min value postion in nums
	// case 1: start < middle < end, search left part
	// case 2: start > middle > end, search right part
	// case 3: start > middle and middle < end  search left part
	// case 4: start < middle and middle > end search  right part

	if start >= end {
		return nums[start]
	}

	middle := (end + start) / 2

	startValue := nums[start]
	endValue := nums[end]
	middleValue := nums[middle]

	if startValue == middleValue && middleValue == endValue {
		return min(findMinRecursive(nums, start, middle), findMinRecursive(nums, middle+1, end))
	} else if startValue <= middleValue && middleValue <= endValue {
		// search left
		return findMinRecursive(nums, start, middle-1)
	} else if startValue > middleValue && middleValue > endValue {
		//search right
		return findMinRecursive(nums, middle, end)
	} else if startValue <= middleValue && middleValue >= endValue {
		// search right
		return findMinRecursive(nums, middle+1, end)
	} else {
		return min(findMinRecursive(nums, start, middle), findMinRecursive(nums, middle+1, end))
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
