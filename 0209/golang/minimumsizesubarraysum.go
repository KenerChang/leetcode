package minimumsizesubarraysum

import (
	// "fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if nums[0] == s {
		return 1
	}

	var now, start int
	var minLen int = math.MaxInt32
	for i, num := range nums {
		// fmt.Printf("i: %d, now: %d, num: %d, sum: %d\n", i, now, num, now+num)
		now += num

		if now < s {
			continue
		}

		// do pop
		for now > s {
			now -= nums[start]
			start++
		}

		if now < s {
			start--
			now += nums[start]
		}

		// shift one back
		minLen = min(i-start+1, minLen)
	}

	if minLen == math.MaxInt32 {
		minLen = 0
	}

	return minLen
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
