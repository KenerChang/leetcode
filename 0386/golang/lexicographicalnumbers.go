package lexicographicalnumbers

func lexicalOrder(n int) []int {
	return recursive(0, n, []int{})
}

func recursive(base, threshold int, nums []int) []int {
	for i := 0; i <= 9; i++ {
		if base+i > threshold {
			return nums
		}

		num := base + i
		if num == 0 {
			continue
		}

		nums = append(nums, num)

		nums = recursive(num*10, threshold, nums)
	}

	return nums
}
