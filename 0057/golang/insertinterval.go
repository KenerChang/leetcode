package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	// This problem can be solved by looking up the position of
	// start point and end point of the newInterval. After finding
	// the position, then we could decide which intervals need to
	// be merged.

	// find position of start point
	startPointPos := findFirstGreaterInterval(intervals, 0, len(intervals), newInterval[0])

	// merge previous interval if overlap
	if startPointPos > 0 && intervals[startPointPos-1][1] >= newInterval[0] {
		startPointPos--
		newInterval[0] = min(newInterval[0], intervals[startPointPos][0])
	}

	endPointPos := findFirstGreaterInterval(intervals, 0, len(intervals), newInterval[1])
	// fmt.Printf("endPointPos: %d\n", endPointPos)

	if endPointPos > 0 {
		newInterval[1] = max(newInterval[1], intervals[endPointPos-1][1])
	}

	// fmt.Printf("start: %d, end: %d, new interval: %v\n", startPointPos, endPointPos, newInterval)

	// create new intervals
	// result := intervals[0:startPointPos]
	result := append([][]int{}, intervals[0:startPointPos]...)

	result = append(result, newInterval)

	if endPointPos <= len(intervals)-1 {
		// fmt.Println("entered")
		result = append(result, intervals[endPointPos:]...)
	}

	return result
}

func findFirstGreaterInterval(intervals [][]int, start, end, point int) int {
	// find the interval index which is the first interval that start
	// point is greater than the point.

	if start == end {
		return start
	}

	middle := (start + end) / 2

	if intervals[middle][0] <= point {
		// go right
		return findFirstGreaterInterval(intervals, middle+1, end, point)
	} else {
		// go left
		return findFirstGreaterInterval(intervals, start, middle, point)
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
