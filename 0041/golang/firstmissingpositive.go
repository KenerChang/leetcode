package firstmissingpositive

func decode(num int) int {
	mask := (1 << 31) - 1

	return num & mask
}

func encode(num int) int {
	// remove negative
	mask := (1 << 31) - 1
	result := num & mask

	mask = (1 << 31)

	return result | mask
}

func firstMissingPositive(nums []int) int {
	// The largest possible missing positive is len(nums) +1, and
	// this happens when all the nums are in rage 1 ~ len(nums).
	// In this case, we can use nums to store if 1 ~ len(nums) are input.
	for i, num := range nums {
		if num < 0 {
			nums[i] = 0
			continue
		}

		decoded := decode(num)
		if decoded > len(nums) || decoded == 0 {
			continue
		}

		numIndex := decoded - 1
		// fmt.Printf("decoded: %d, encoded: %d\n", decoded, encode(numIndex))
		nums[numIndex] = encode(nums[numIndex])
	}

	// fmt.Printf("nums: %v\n", nums)

	for i, num := range nums {
		if num < 0 {
			return i + 1
		}

		mask := (1 << 31)
		exist := num & mask
		// fmt.Printf("exist: %d\n", exist)
		if exist == 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
