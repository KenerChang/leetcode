package containsduplicateiii

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// hashMap := map[int]int{}
	for i, num := range nums {
		for j := 1; j <= k && (i+j) < len(nums); j++ {
			diff := nums[i+j] - num
			if diff <= t && diff >= -t {
				return true
			}
		}
	}
	return false
}
