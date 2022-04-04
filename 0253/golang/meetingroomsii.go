package meetingroomsii

import (
	"container/heap"
	"sort"
)

type mIntervals [][]int

func (intervals mIntervals) Len() int {
	return len(intervals)
}

func (intervals mIntervals) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func (intervals mIntervals) Less(i, j int) bool { return intervals[i][0] <= intervals[j][0] }

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	// return 1 if only one interval
	if len(intervals) == 1 {
		return 1
	}

	// sort intervals by start time
	sort.Sort(mIntervals(intervals))

	endTimes := &IntHeap{}
	heap.Init(endTimes)
	neededRoom := 0
	// for interval in intervals
	for _, interval := range intervals {
		// use heap to keep min end time
		if endTimes.Len() == 0 {
			heap.Push(endTimes, interval[1])
		} else {
			// keep pop out intervals whichs end time less then current interval start time
			minEndTime := heap.Pop(endTimes).(int)
			for minEndTime <= interval[0] && endTimes.Len() > 0 {
				minEndTime = heap.Pop(endTimes).(int)
			}

			if minEndTime > interval[0] {
				heap.Push(endTimes, minEndTime)
			}
			heap.Push(endTimes, interval[1])

		}

		// check if needed rooms need to update
		neededRoom = max(neededRoom, endTimes.Len())
	}

	return neededRoom
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
