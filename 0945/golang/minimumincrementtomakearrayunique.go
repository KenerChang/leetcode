package minimumincrementtomakearrayunique

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)

	var minNonConfictNum int = nums[0]
	var incrementCount int
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if num > minNonConfictNum {
			minNonConfictNum = num
		} else {
			incrementCount += (minNonConfictNum - num) + 1
			minNonConfictNum++
		}
	}

	return incrementCount
}
