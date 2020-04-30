package uniquebinarysearchtrees

import (
	"strconv"
)

func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	cache := map[string]int{}
	return numTreesImpl(1, n, cache)
}

func numTreesImpl(start, end int, cache map[string]int) (num int) {
	key := toKey(start, end)
	if num, found := cache[key]; found {
		return num
	}

	if start == end {
		return 1
	}

	for idx := start; idx <= end; idx++ {
		if idx == start {
			num += numTreesImpl(start+1, end, cache)
		} else if idx == end {
			num += numTreesImpl(start, end-1, cache)
		} else {
			leftTreeNum := numTreesImpl(start, idx-1, cache)
			rightTreeNum := numTreesImpl(idx+1, end, cache)
			num += (leftTreeNum * rightTreeNum)
		}
	}
	cache[key] = num

	return
}

func toKey(start, end int) string {
	return strconv.Itoa(start) + "," + strconv.Itoa(end)
}
