package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = 1
	}

	var left int = 1
	var right int = 1
	for i := 0; i < len(nums); i++ {
		result[i] *= left

		j := len(nums) - 1 - i
		result[j] *= right

		left *= nums[i]
		right *= nums[j]
	}

	return result
}
