package splitarrayintoconsecutivesubsequences

func isPossible(nums []int) bool {
	// To solve this problem, we need iterate nums.
	// For each num, we have 2 situations:
	// 1. add into an existed subsequence
	// 2. start a new subsequence

	// To handle these 2 situations, we need two maps for keeping
	// 1. frequency of each nunmber
	// 2. subsequence count ending with each num

	frequencies := map[int]int{}
	for _, num := range nums {
		frequencies[num]++
	}

	subsequencies := map[int]int{}
	for _, num := range nums {
		if frequencies[num] == 0 {
			// has been used
			continue
		}

		if subsequencies[num-1] == 0 {
			// can not append to an existed subsequencies
			if frequencies[num+1] <= 0 || frequencies[num+2] <= 0 {
				// impossible to form a subsequence
				return false
			}

			subsequencies[num+2]++
			frequencies[num]--
			frequencies[num+1]--
			frequencies[num+2]--
		} else {
			subsequencies[num-1]--
			subsequencies[num]++
			frequencies[num]--
		}
	}

	return true
}
