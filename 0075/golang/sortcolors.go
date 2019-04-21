package sortcolors

// sortColors sorts colors (which are represent in numbers) in place
// Whenever sortColors meets 0 or 2, it make sure all 0s and 2s are sorted
func sortColors(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	// red: 0, white: 1, and blue: 2
	var temp int
	zeroIdx := 0            // where last zero should be put
	twoIdx := len(nums) - 1 // where last two should be put

	for idx, num := range nums {
		if num == 0 {
			// put the zero where to the last 0 and update zero pointer
			if idx > zeroIdx {
				for zeroIdx < idx && nums[zeroIdx] == 0 {
					zeroIdx++
				}

				temp = nums[zeroIdx]
				nums[idx] = temp
				nums[zeroIdx] = num

				if temp == 2 {
					// if we switch a two, make sure current 2s are sorted
					for twoIdx > idx && nums[twoIdx] == 2 {
						twoIdx--
					}

					temp = nums[twoIdx]
					nums[idx] = temp
					nums[twoIdx] = 2
				}
			}
		}

		if num == 2 {
			// same as process of meeting a zero in reverse way
			if twoIdx > idx {
				for twoIdx > idx && nums[twoIdx] == 2 {
					twoIdx--
				}

				temp = nums[twoIdx]
				nums[idx] = temp
				nums[twoIdx] = num

				if temp == 0 {
					for zeroIdx < idx && nums[zeroIdx] == 0 {
						zeroIdx++
					}

					temp = nums[zeroIdx]
					nums[idx] = temp
					nums[zeroIdx] = 0
				}
			}
		}
	}
}
