package insertionsortlist

import (
	"testing"
)

func verify(head *ListNode) ([]int, bool) {
	if head == nil {
		return []int{}, true
	}

	sorted := true
	list := []int{}
	node := head
	for node != nil {
		if len(list) > 0 && node.Val < list[len(list)-1] {
			sorted = false
		}

		list = append(list, node.Val)
		node = node.Next
	}
	return list, sorted
}

func TestInsertionSortList(t *testing.T) {
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

	newHead := insertionSortList(head)
	if list, sorted := verify(newHead); !sorted {
		t.Errorf("got: %+v", list)
	}
}

func TestInsertionSortListII(t *testing.T) {
	head := &ListNode{
		Val: 4,
	}

	newHead := insertionSortList(head)
	if list, sorted := verify(newHead); !sorted {
		t.Errorf("got: %+v", list)
	}
}

func TestInsertionSortListIII(t *testing.T) {
	newHead := insertionSortList(nil)
	if list, sorted := verify(newHead); !sorted {
		t.Errorf("got: %+v", list)
	}
}

func TestInsertionSortListIV(t *testing.T) {
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

	newHead := insertionSortList(head)
	if list, sorted := verify(newHead); !sorted {
		t.Errorf("got: %+v", list)
	}
}

func TestInsertionSortListV(t *testing.T) {
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

	newHead := insertionSortList(head)
	if list, sorted := verify(newHead); !sorted {
		t.Errorf("got: %+v", list)
	}
}
