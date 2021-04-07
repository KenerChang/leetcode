package majorityelementii

import (
	"math"
)

func majorityElement(nums []int) []int {
	counts := map[int]int{}
	results := []int{}
	total := len(nums)
	threshold := int64(math.Floor(float64(total) / 3))
	inResult := map[int]bool{}

	for _, num := range nums {
		var count int
		var ok bool
		if count, ok = counts[num]; ok {
			count = count + 1
		} else {
			count = 1
		}
		counts[num] = count

		if _, ok = inResult[num]; count > int(threshold) && !ok {
			results = append(results, num)
			inResult[num] = true
		}
	}

	return results
}
