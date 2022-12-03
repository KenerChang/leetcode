package reorganizestring

import (
	"container/heap"
)

type char struct {
	b    byte
	freq int
}

type charHeap []char

func (h charHeap) Len() int { return len(h) }

func (h charHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }

func (h charHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *charHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(char))
}

func (h *charHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func reorganizeString(s string) string {
	// we can use a max heap to solve this problem
	charMap := map[byte]int{}
	for _, c := range s {
		charMap[byte(c)]++
	}

	chars := []char{}
	for k, v := range charMap {
		chars = append(chars, char{b: k, freq: v})
	}
	h := charHeap(chars)
	heap.Init(&h)

	var result []byte

	for h.Len() > 0 {
		// first pop the most frequency char from heap (a)
		mostFreq := heap.Pop(&h).(char)

		// fmt.Printf("mostFreq: %s\n", string(mostFreq.b))

		if len(result) == 0 || result[len(result)-1] != mostFreq.b {
			result = append(result, mostFreq.b)
			mostFreq.freq--
		} else {
			if h.Len() == 0 {
				return ""
			}

			nextMostFreq := heap.Pop(&h).(char)
			// fmt.Printf("nextMostFreq: %s\n", string(nextMostFreq.b))
			result = append(result, nextMostFreq.b)
			nextMostFreq.freq--

			if nextMostFreq.freq > 0 {
				heap.Push(&h, nextMostFreq)
			}
		}

		if mostFreq.freq > 0 {
			heap.Push(&h, mostFreq)
		}
	}
	return string(result)
}
