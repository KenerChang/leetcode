package jumpgamevi

import (
	"container/heap"
)

type maxHeap [][]int // 0: pos 1: value

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}

func (h *maxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxResult(nums []int, k int) int {
	// we can use dp & min heap to solve this problem
	dp := make([]int, len(nums))

	h := &maxHeap{}
	heap.Init(h)

	for i := range nums {
		if h.Len() > 0 {
			// pop top until it is in range
			// fmt.Printf("h: %v\n", h)
			top := heap.Pop(h)
			for i-top.([]int)[0] > k {
				top = heap.Pop(h)
			}

			// fmt.Printf("i:%d, choose: %v\n", i, top)
			dp[i] = top.([]int)[1] + nums[i]
			heap.Push(h, top)
		} else {
			dp[i] = nums[i]
		}

		heap.Push(h, []int{i, dp[i]})
	}

	// fmt.Printf("dp: %v\n", dp)
	return dp[len(dp)-1]
}
