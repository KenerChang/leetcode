package findtheduplicatenumber

func findDuplicate(nums []int) int {
	// We can reduce the problem to find the intersection point of a cycled linked list
	// and we can solve this problem by using Floyd's algorithm

	if len(nums) == 2 {
		return 1
	}

	// define two pointers: fast and slow
	var fast, slow int = 0, 0

	// phase1: move pointers until they meet
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]

		if fast == slow {
			break
		}
	}

	// phase2: move pointers again until they meet again, the point is the intersection point
	// put slow to start point
	slow = 0
	for {
		fast = nums[fast]
		slow = nums[slow]

		if fast == slow {
			break
		}
	}

	return fast
}
