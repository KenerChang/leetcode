package mergeinterval

import "sort"

type itvals [][]int // intervals

func (is itvals) Len() int {
	return len(is)
}

func (is itvals) Less(i, j int) bool {
	return is[i][0] < is[j][0]
}

func (is itvals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	// sort intervals by start
	// then iterate intervals if itervals[i][0] <= intervals[i-1][1], then merge the two intervals

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Sort(itvals(intervals))

	current := intervals[0]
	results := [][]int{}
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[0] > current[1] {
			results = append(results, current)
			// no overlap
			current = interval
		} else {
			current[1] = max(current[1], interval[1])
		}
	}

	results = append(results, current)

	return results
}
