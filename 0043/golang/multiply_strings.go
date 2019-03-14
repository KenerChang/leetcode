package multiply_strings

import (
	"strconv"
)

func multiply(mulitplican string, multiplier string) string {
	// if mulitplier or Multiplican is zero
	// return 0
	if mulitplican == "0" || multiplier == "0" {
		return "0"
	}

	// transform mulitplican and multiplier to slice of int8
	nums1 := toSlice(mulitplican)
	nums2 := toSlice(multiplier)

	// because result of n digits x m digits is bounded by m+n digits
	digits := len(nums1) + len(nums2)
	nums := make([]int, digits) 

	// fill up each digits in nums
	for idx1, num1 := range nums1 {
		for idx2, num2 := range nums2 {
			total := num1 * num2
			od := total % 10 // digits in ones
			td := total / 10 // digits in tens

			odIdx := idx1 + idx2
			nums[odIdx] += od
			// check if carry
			if nums[odIdx] >= 10 {
				td++
				nums[odIdx] -= 10
			}

			tdIdx := idx1 + idx2 + 1
			nums[tdIdx] += td
			// check if carray
			if nums[tdIdx] >= 10 {
				nums[tdIdx+1]++
				nums[tdIdx] -= 10
			}
		}
	}
	return toString(nums)
}

func toSlice(numStr string) []int {
	intSlice := make([]int, len(numStr))
	idx := 0
	base := 48 // index of 0 in ascii
	for i:=len(numStr) - 1; i>= 0; i-- {
		intSlice[idx] = int(numStr[i]) - base
		idx++
	}
	return intSlice
}

func toString(nums []int) string {
	result := ""
	for i := len(nums) - 1; i >=0; i-- {
		result += strconv.Itoa(nums[i])
	}
	
	// remove first zero
	if int(result[0]) == 48 {
		result = result[1:]
	}
	return result
}