package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{fmt.Sprint(nums[0])}
	}

	size := len(nums)
	start := nums[0]
	shift := 1
	results := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[i]-start != shift {
			if shift == 1 {
				results = append(results, fmt.Sprint(start))
			} else {
				results = append(results, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			shift = 1
			start = nums[i]
		} else {
			shift++
		}
	}

	if nums[size-1]-start != shift {
		if shift == 1 {
			results = append(results, fmt.Sprint(start))
		} else {
			results = append(results, fmt.Sprintf("%d->%d", start, nums[size-1]))
		}
	}
	return results
}
