package majorityelement

func majorityElement(nums []int) int {
	counts := map[int]int{}
	threshold := len(nums) / 2
	for _, num := range nums {
		count, found := counts[num]
		if found {
			count++
		} else {
			count = 1
		}

		if count > threshold {
			return num
		}
		counts[num] = count
	}
	return 0
}
