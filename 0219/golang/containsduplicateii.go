package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	for i, num := range nums {
		for j := i + 1; (j-i <= k) && (j < len(nums)); j++ {
			if nums[j] == num {
				return true
			}
		}
	}
	return false
}
