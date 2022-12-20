package maximumaveragesubarrayi

func findMaxAverage(nums []int, k int) float64 {
	var largest int
	var ksum int

	for i := 0; i < k; i++ {
		ksum += nums[i]
	}
	largest = ksum

	for i := k; i < len(nums); i++ {
		ksum -= nums[i-k]
		ksum += nums[i]

		largest = max(largest, ksum)
	}

	return float64(largest) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
