package longestincreasingsubsequence

func firstGTE(sorted []int, l, r int, target int) int {
	if l >= r {
		return l
	}

	mid := (l + r) / 2

	if sorted[mid] >= target {
		// search left
		return firstGTE(sorted, l, mid, target)
	} else {
		//  search right
		return firstGTE(sorted, mid+1, r, target)
	}
}

func lengthOfLIS(nums []int) int {
	// We can use an array to keep number list
	// for each num, we have 3 situations
	// 1. the num is less than least number in the list, replace first num
	// 2. the num is larger than largest number in the list, append to the list
	// 3. the num is between the least and the largest, find a number which is the first number larger than the num, and replace it

	if len(nums) <= 1 {
		return len(nums)
	}

	list := make([]int, len(nums)+1)
	list[0] = nums[0]
	listP := 1

	for _, num := range nums {
		if num < list[0] {
			list[0] = num
		} else if num > list[listP-1] {
			list[listP] = num
			listP++
		} else {
			pos := firstGTE(list, 0, listP-1, num)
			list[pos] = num
		}
	}

	return listP
}
