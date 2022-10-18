package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	// Consider that if sum of sum[i] - sum[j] is 0, it means that sum of nums between nums[i] and nunms[j] is 0.
	// Extend this thought, if sums[j] - sum[i] is k, it means nums[i+1] to nums[j] is a subarray which sums to k
	// so we can use a map to keep sums of nums and for every num, we get sum of num[i], and see if sum[i] - k is in map

	countMap := map[int]int{
		0: 1, // for sum[i] = k
	}
	var result int
	var sum int
	for _, num := range nums {
		sum += num

		// fmt.Printf("sum: %d, sum-k: %d, countMap: %v\n", sum, sum-k, countMap)
		if count, ok := countMap[sum-k]; ok {
			result += count
		}

		countMap[sum]++
	}

	return result
}
