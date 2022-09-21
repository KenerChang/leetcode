package golang

import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	minHeap minHeap
	maxHeap maxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: make(minHeap, 0),
		maxHeap: make(maxHeap, 0),
	}
}

func (m *MedianFinder) AddNum(num int) {
	// put num to max heap
	heap.Push(&m.maxHeap, num)

	// pop max num from max heap to min heap
	heap.Push(&m.minHeap, heap.Pop(&m.maxHeap))

	// reblance if needed
	if m.minHeap.Len() > m.maxHeap.Len() {
		heap.Push(&m.maxHeap, heap.Pop(&m.minHeap))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	var result float64
	if len(m.maxHeap) > len(m.minHeap) {
		// odd
		result = float64(m.maxHeap[0])
	} else {
		// even
		result = float64(m.maxHeap[0]+m.minHeap[0]) / 2.0
	}
	return result
}
