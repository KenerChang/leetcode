package randompickindex

import "math/rand"

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := map[int][]int{}

	for i, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = []int{i}
		} else {
			m[num] = append(m[num], i)
		}
	}
	return Solution{m: m}
}

func (this *Solution) Pick(target int) int {
	return this.m[target][rand.Intn(len(this.m[target]))]
}
