package topkfrequentelements

import "sort"

type Number struct {
	Val  int
	Freq int
}

type Numbers []*Number

func (ns Numbers) Len() int {
	return len(ns)
}

func (ns Numbers) Less(i, j int) bool {
	return ns[i].Freq > ns[j].Freq
}

func (ns Numbers) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func topKFrequent(nums []int, k int) []int {
	numMap := map[int]*Number{}

	numbers := []*Number{}
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			numMap[num].Freq++
		} else {
			number := &Number{Val: num, Freq: 1}
			numMap[num] = number
			numbers = append(numbers, number)
		}
	}

	sort.Sort(Numbers(numbers))

	results := []int{}
	for i := 0; i < k; i++ {
		results = append(results, numbers[i].Val)
	}
	return results
}
