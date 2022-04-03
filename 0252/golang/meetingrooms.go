package meetingrooms

import (
	"sort"
)

type mIntervals [][]int

// Len is part of sort.Interface.
func (internvals mIntervals) Len() int {
	return len(internvals)
}

// Swap is part of sort.Interface.
func (internvals mIntervals) Swap(i, j int) {
	internvals[i], internvals[j] = internvals[j], internvals[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (intervals mIntervals) Less(i, j int) bool {
	return intervals[i][0] <= intervals[j][0]
}

func canAttendMeetings(intervals [][]int) bool {
	// if only one interval, return true
	if len(intervals) == 1 || len(intervals) == 0 {
		return true
	}

	// sort intervals by their start time
	sort.Sort(mIntervals(intervals))

	// for each iterval, we keep max end time of preivous intervals
	// if start time of the interval is less than max end time of previous intervals
	// there must have conflict
	maxEndTime := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < maxEndTime {
			return false
		}

		maxEndTime = max(maxEndTime, intervals[i][1])
	}

	return true
}

func max(a, b int) int {
	if a <= b {
		return b
	}

	return a
}
