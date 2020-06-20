package findminimuminrotatedsortedarray

func findMin(nums []int) int {
	last := len(nums) - 1
	middle := len(nums) / 2

	if last == 0 || (nums[0] <= nums[middle] && nums[middle] <= nums[last]) {
		return nums[0]
	}

	if len(nums) == 2 {
		return min(nums[0], nums[1])
	}

	if nums[middle-1] > nums[middle] && nums[middle+1] > nums[middle] {
		return nums[middle]
	}

	if nums[0] > nums[middle] {
		return findMin(nums[0:middle])
	}
	return findMin(nums[middle+1:])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
