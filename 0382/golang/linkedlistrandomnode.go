package linkedlistrandomnode

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		Head: head,
	}
}

func (s *Solution) GetRandom() int {
	// Use reservoir sampling
	scope := float64(1)

	n := s.Head
	var result int
	for n != nil {
		if rand.Float64() < float64(1)/scope {
			result = n.Val
		}

		n = n.Next
		scope += 1
	}
	return result
}
