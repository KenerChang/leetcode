package sortlist

import (
	"testing"
)

func verify(head *ListNode, target []int) ([]int, bool) {
	list := []int{}
	node := head
	for node != nil {
		list = append(list, node.Val)
		node = node.Next
	}

	if len(list) != len(target) {
		return list, false
	}

	for i, numt := range target {
		if list[i] != numt {
			return list, false
		}
	}

	return list, true
}

func TestSortList(t *testing.T) {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	target := []int{1, 2, 3, 4}

	newHead := sortList(head)
	if list, sorted := verify(newHead, target); !sorted {
		t.Errorf("expect: %+v, got: %+v", target, list)
	}
}

func TestSortListII(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	target := []int{1}

	newHead := sortList(head)
	if list, sorted := verify(newHead, target); !sorted {
		t.Errorf("expect: %+v, got: %+v", target, list)
	}
}

func TestSortListIII(t *testing.T) {
	newHead := sortList(nil)
	target := []int{}
	if list, sorted := verify(newHead, target); !sorted {
		t.Errorf("expect: %+v, got: %+v", target, list)
	}
}

func TestSortListIV(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	target := []int{1, 2, 3, 4}

	newHead := sortList(head)
	if list, sorted := verify(newHead, target); !sorted {
		t.Errorf("expect: %+v, got: %+v", target, list)
	}
}

func TestSortListV(t *testing.T) {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	target := []int{1, 2, 3, 4}

	newHead := sortList(head)
	if list, sorted := verify(newHead, target); !sorted {
		t.Errorf("expect: %+v, got: %+v", target, list)
	}
}
