package rotatelist

import (
	"testing"
)

func TestRotateRight(t *testing.T) {
	listHead := &ListNode{
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
	target := []int{4, 5, 1, 2, 3}

	resultList := rotateRight(listHead, 2)
	if !isSame(resultList, target) {
		t.Errorf("expect %+v, but got: %+v", target, resultList)
	}

	listHead = &ListNode{
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
	target = []int{5, 1, 2, 3, 4}

	resultList = rotateRight(listHead, 1)
	if !isSame(resultList, target) {
		t.Errorf("expect %+v, but got: %+v", target, resultList)
	}

	listHead = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	target = []int{2, 0, 1}

	resultList = rotateRight(listHead, 4)
	if !isSame(resultList, target) {
		t.Errorf("expect %+v, but got: %+v", target, resultList)
	}
}

func isSame(node *ListNode, target []int) bool {
	idx := 0
	cNode := node

	result := true
	for cNode != nil {
		if idx >= len(target) {
			result = false
			break
		}

		if cNode.Val != target[idx] {
			result = false
			break
		}
		cNode = cNode.Next
		idx++
	}

	if idx != len(target) {
		result = false
	}

	return result
}
