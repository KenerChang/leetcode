package mergeksortedlists

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	// check if lists is empty
	if len(lists) == 0 {
		return nil
	}

	// init heap
	h := &ListNodeHeap{}
	heap.Init(h)
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	var head, tail *ListNode
	for h.Len() > 0 {
		// every time we put smallest element of each list into heap
		// get the smallest elemt from the heap
		// put it to the end of result list
		node := heap.Pop(h).(*ListNode)
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = tail.Next
		}

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return head
}
