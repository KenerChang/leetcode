package reverselinkedlistii

import (
	"testing"
)

func verify(list *ListNode, target []int) bool {
	if list == nil {
		if len(target) > 0 {
			return false
		}
		return true
	}

	idx := 0
	curNode := list
	for curNode != nil {
		if idx >= len(target) {
			return false
		}

		if curNode.Val != target[idx] {
			return false
		}
		idx++
		curNode = curNode.Next
	}
	return true
}

func TestReverseBetween1(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	list = reverseBetween(list, 2, 4)
	target := []int{1, 4, 3, 2, 5}
	if !verify(list, target) {
		t.Errorf("failed, target: %+v got: %+v", target, list.Next.Val)
	}
}

func TestReverseBetween2(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	list = reverseBetween(list, 1, 7)
	target := []int{7, 6, 5, 4, 3, 2, 1}
	if !verify(list, target) {
		t.Errorf("failed, target: %+v got: %+v", target, list)
	}
}

func TestReverseBetweenSameStartAndEnd(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	list = reverseBetween(list, 1, 1)
	target := []int{1, 2, 3, 4, 5, 6, 7}
	if !verify(list, target) {
		t.Errorf("failed, target: %+v got: %+v", target, list)
	}

	list = reverseBetween(list, 7, 7)
	if !verify(list, target) {
		t.Errorf("failed, target: %+v got: %+v", target, list)
	}
}

func TestReverseBetweenStartLargerThanEnd(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	list = reverseBetween(list, 5, 4)
	target := []int{1, 2, 3, 4, 5, 6, 7}
	if !verify(list, target) {
		t.Errorf("failed, target: %+v got: %+v", target, list)
	}
}
