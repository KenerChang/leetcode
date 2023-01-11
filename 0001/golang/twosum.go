package twosum

func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{} // store num & its position mapping

	for i, num := range nums {
		complement := target - num

		if cidx, ok := numsMap[complement]; ok {
			return []int{cidx, i}
		}

		numsMap[num] = i
	}

	return []int{}
}
