package superuglynumber

import (
	"container/heap"
)

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

func nthSuperUglyNumber(n int, primes []int) int {
	// we can generate all numbers from these primes until n is reached
	if n == 1 {
		return 1
	}

	// use a heap which help us to get current mininum number
	intHeap := &IntHeap{}
	heap.Init(intHeap)

	for _, prime := range primes {
		heap.Push(intHeap, prime)
	}

	// // use a dictionary to filter duplicated numbers
	// seen := map[int]struct{}{}

	count := 1
	var current int
	var last int = 1
	for count < n {
		current = heap.Pop(intHeap).(int)

		if last == current {
			continue
		}
		last = current
		count++

		// fmt.Printf("current: %d, count: %d\n", current, count)

		// put all the possible numbers to heap
		for _, prime := range primes {
			heap.Push(intHeap, current*prime)
		}
	}
	return current
}
