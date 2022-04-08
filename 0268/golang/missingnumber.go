package missingnumber

func missingNumber(nums []int) int {
	// count vault trapezoid formula
	count := trapezoid(len(nums))

	// for num in nums
	for _, num := range nums {
		count -= num
	}

	// return count
	return count
}

func trapezoid(limit int) int {
	return (1 + limit) * limit / 2
}
