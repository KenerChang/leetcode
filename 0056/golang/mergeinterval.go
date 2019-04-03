package mergeinterval

import "sort"

type Interval struct {
	Start int
	End   int
}

type ByStart []Interval

func (is ByStart) Len() int           { return len(is) }
func (is ByStart) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
func (is ByStart) Less(i, j int) bool { return is[i].Start < is[j].Start }

func merge(intervals []Interval) []Interval {
	// first sort the intervals by Start
	// with each interval in intervals
	// we just need to check if it should be merge with previous one
	// the time complexity is O(nlogn)

	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	sort.Sort(ByStart(intervals))
	result := []Interval{}
	for idx, interval := range intervals {
		if idx == 0 {
			result = append(result, interval)
			continue
		}

		previous := &result[len(result)-1]
		if interval.Start >= previous.Start && interval.Start <= previous.End {
			previous.End = Max(interval.End, previous.End)
		} else {
			result = append(result, interval)
		}
	}
	return result
}

func Max(i, j int) int {
	if i <= j {
		return j
	} else {
		return i
	}
}
