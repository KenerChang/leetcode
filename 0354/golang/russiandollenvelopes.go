package russiandollenvelopes

import (
	"sort"
)

type byWidthAndHeight [][]int

func (b byWidthAndHeight) Len() int {
	return len(b)
}

func (b byWidthAndHeight) Less(i, j int) bool {
	if b[i][0] != b[j][0] {
		return b[i][0] < b[j][0]
	}

	return b[i][1] > b[j][1]
}

func (b byWidthAndHeight) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func firstGTE(sorted []int, l, r, target int) int {
	if l >= r {
		return l
	}

	mid := (l + r) / 2
	if sorted[mid] == target {
		return mid
	} else if sorted[mid] > target {
		return firstGTE(sorted, l, mid, target)
	} else {
		return firstGTE(sorted, mid+1, r, target)
	}
}

func lis(nums []int) int {
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

func maxEnvelopes(envelopes [][]int) int {
	// This is a 2-dimension LIS problem, the question is how we solve this with LIS
	// To reduce this problem into a 1-dimension LIS problem
	// we can sort envelopes in width and height in reverse order
	// so we can prevent envelopes with same width apear in the LIS

	sort.Sort(byWidthAndHeight(envelopes))

	heights := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		heights[i] = envelopes[i][1]
	}

	return lis(heights)
}
