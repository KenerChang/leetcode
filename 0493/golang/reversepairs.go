package reversepairs

import (
	"sort"
)

func reversePairs(nums []int) int {
	// We can solve this problem by using Binary Indexed Tree (BIT).
	// To use BIT, we need to reduce the problem as a prefix sum problem.
	// We can reduce reverse paris problem to prefix sum problem by these steps:
	// 1. we need a sorted array so that we can know how many numbers are smaller than math.Floor(num[i]/2)
	// 2. reverse traverse the array, so that only numbers after nums[i] will show in BIT
	// in each iteration, we query the BIT to get the count of numbers which are smaller than math.Floor(num[i]/2)
	// after that, we update the BIT by adding nums[i] into it

	// delcare a BIT, and a sorted array
	bit := make([]int, len(nums)+1)
	sorted := make([]int, len(nums))
	copy(sorted, nums)

	// sort
	sort.Ints(sorted)

	// reverse traverse nums
	var result int
	for _, num := range nums {
		// query the BIT
		// fmt.Printf("nums[%d]=%d, sorted=%v, bit=%v, nums[%d]=%d\n", j, nums, sorted, bit, j, nums[j])
		target := 2*num + 1

		targetIdxInBIT := index(sorted, target)
		result += query(bit, targetIdxInBIT)

		// update the BIT
		update(bit, index(sorted, num))

	}

	return result
}

func index(sorted []int, num int) int {
	// find the last index of num in sorted array
	l, r := 0, len(sorted)-1
	for l <= r {
		mid := (l + r) / 2
		if sorted[mid] >= num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l + 1
}

func query(bit []int, idx int) int {
	var result int
	for idx < len(bit) {
		result += bit[idx]
		idx += idx & -idx
	}

	return result
}

func update(bit []int, idx int) {
	for idx > 0 {
		bit[idx]++
		idx -= idx & -idx
	}
}
