package kthlargestelementinanarray

import (
// "fmt"
// "sort"
)

func findKthLargest(nums []int, k int) int {
	return kSelect(nums, 0, len(nums)-1, k)
}

func kSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	boundary := right
	lp, rp := left, right-1

	for lp <= rp {
		for lp <= rp && nums[lp] >= nums[boundary] {
			lp++
		}

		for lp <= rp && nums[rp] < nums[boundary] {
			rp--
		}

		if lp <= rp {
			nums[lp], nums[rp] = nums[rp], nums[lp]
			lp++
			rp--
		}
	}

	nums[lp], nums[boundary] = nums[boundary], nums[lp]
	boundary = lp

	if boundary > k-1 {
		return kSelect(nums, left, boundary-1, k)
	} else if boundary < k-1 {
		return kSelect(nums, boundary+1, right, k)
	}

	return nums[boundary]
}
