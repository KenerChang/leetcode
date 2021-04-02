package containsduplicateiii

func getBucket(num, size int) int {
	if num >= 0 {
		return num / size
	}
	return -(-num-1)/size - 1
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	buckets := make(map[int]int)
	bucketSize := max(t, 1)

	for i, v := range nums {
		bucketIndex := getBucket(v, bucketSize)
		if _, ok := buckets[bucketIndex]; ok {
			return true
		}
		if k != 0 {
			if prev, ok := buckets[bucketIndex-1]; ok && v-prev <= t {
				return true
			}
			if next, ok := buckets[bucketIndex+1]; ok && next-v <= t {
				return true
			}
		}

		if i >= k {
			delete(buckets, getBucket(nums[i-k], bucketSize))
		}
		buckets[bucketIndex] = v
	}
	return false
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
