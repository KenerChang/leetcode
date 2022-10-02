package maxstack

import (
	"container/heap"
	"container/list"
)

type MaxHeap []*list.Element

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return h[i].Value.(int) >= h[j].Value.(int)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	item := x.(*list.Element)
	*h = append(*h, item)
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// We can implement MaxStack with a double linked list and a max heap.
type MaxStack struct {
	l       *list.List
	h       MaxHeap
	removed map[*list.Element]bool
}

func Constructor() MaxStack {
	return MaxStack{
		h:       []*list.Element{},
		l:       list.New(),
		removed: map[*list.Element]bool{},
	}
}

func (s *MaxStack) Push(x int) {
	// add it to the heap, and add it to the tail of the linked list
	ele := s.l.PushBack(x)

	heap.Push(&s.h, ele)
}

func (s *MaxStack) Pop() int {
	// remove the last element in the list, and also remove it from the heap
	ele := s.l.Back()
	s.l.Remove(ele)
	s.removed[ele] = true

	return ele.Value.(int)
}

func (s *MaxStack) Top() int {
	// return the last element in the list
	ele := s.l.Back()
	return ele.Value.(int)
}

func (s *MaxStack) PeekMax() int {
	// get the max element from the heap
	maxEle := s.h[0]
	for s.removed[maxEle] {
		s.removed[maxEle] = false
		heap.Pop(&s.h)
		maxEle = s.h[0]
	}
	return maxEle.Value.(int)
}

func (s *MaxStack) PopMax() int {
	// get the max element from the heap, and remove it from the list

	maxEle := heap.Pop(&s.h).(*list.Element)
	for s.removed[maxEle] {
		s.removed[maxEle] = false
		maxEle = heap.Pop(&s.h).(*list.Element)
	}

	s.l.Remove(maxEle)

	return maxEle.Value.(int)
}
