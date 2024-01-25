package movezeroes

func moveZeroes(nums []int) {
	var non_zero_count int

	for i, num := range nums {
		if num != 0 {
			// swap and increase count
			nums[i] = 0
			nums[non_zero_count] = num
			non_zero_count++
		}
	}
}
