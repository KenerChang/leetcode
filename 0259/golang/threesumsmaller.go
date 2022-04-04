package threesumsmaller

import (
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
	// if nums length is less than 3, return 0
	if len(nums) < 3 {
		return 0
	}

	// sort nums
	sort.Ints(nums)

	result := 0
	// for num in nums
	for i := 0; i < len(nums)-1; i++ {
		// result += nSumSmaller(nums[i+1:], target - num, 2)
		result += twoSumSmaller(nums[i+1:], target-nums[i])
		// fmt.Printf("num: %d, result: %d\n", nums[i], result)
	}
	return result
}

func twoSumSmaller(sortedNums []int, target int) int {
	// use two ptr

	// if sortedNums[ptrHead] + sortedNums[ptrTail] < target, ptrHead++, result += ptrTail - ptrHead
	// else ptrTrail--

	// do this until ptrHead >= ptrTail

	ptrHead := 0
	ptrTail := len(sortedNums) - 1
	result := 0
	for ptrHead < ptrTail {
		if sortedNums[ptrHead]+sortedNums[ptrTail] < target {
			result += ptrTail - ptrHead
			ptrHead++
		} else {
			ptrTail--
		}
	}
	return result
}
