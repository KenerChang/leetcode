package hindex

import (
	"sort"
)

func hIndex(citations []int) int {
	// sort citations desc
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	most := min(citations[0], len(citations))
	// from most citation number to least citation number, find max h
	i := 0
	for h := most; h > 0; h-- {
		// fmt.Printf("h: %d\n", h)
		// if i+1 >= h, return
		if i >= h {
			return h
		}

		// move i to next val
		for i < len(citations) && citations[i] >= h {
			i++
		}

		// fmt.Printf("i: %d\n", i)
		// if accumulated citations is larger than h, return h
		if i >= h {
			return h
		}
	}
	return 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
