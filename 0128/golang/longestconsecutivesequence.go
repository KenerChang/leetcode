package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	// use a map to store number
	numsMap := map[int]bool{}
	for _, num := range nums {
		numsMap[num] = true
	}

	// for num in nums, find all n that n -1 is not in nums
	// because these numbers might be start of the consecutive sequence
	longestLength := 0
	for _, num := range nums {
		if _, ok := numsMap[num-1]; !ok {
			currentLength := 1
			next := num + 1

			for {
				if _, ok := numsMap[next]; ok {
					currentLength++
					next++
				} else {
					break
				}
			}
			longestLength = max(longestLength, currentLength)
		}
	}
	return longestLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
