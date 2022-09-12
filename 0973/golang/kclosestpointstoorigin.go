package kclosestpointstoorigin

import (
	"container/heap"
)

type Points []Point

func (h Points) Len() int { return len(h) }
func (h Points) Less(i, j int) bool {
	pointI := h[i]
	pointJ := h[j]

	return (pointI.X*pointI.X + pointI.Y*pointI.Y) < (pointJ.X*pointJ.X + pointJ.Y*pointJ.Y)
}
func (h Points) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Points) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Point))
}

func (h *Points) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Point struct {
	X int
	Y int
}

func kClosest(points [][]int, k int) [][]int {
	// we can use heap to solve this problem
	// and the time complexity is O(klogn) + O(n)

	// init heap
	h := &Points{}
	heap.Init(h)
	for _, point := range points {
		p := Point{X: point[0], Y: point[1]}
		heap.Push(h, p)
	}

	// pop k elements
	results := [][]int{}
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(Point)

		results = append(results, []int{p.X, p.Y})
	}

	return results
}
