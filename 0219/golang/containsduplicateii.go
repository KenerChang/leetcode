package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	hashMap := map[int]int{}
	for i, num := range nums {
		if prev, ok := hashMap[num]; ok {
			if i-prev <= k {
				return true
			}
		}
		hashMap[num] = i
	}
	return false
}
