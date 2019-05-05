package removeduplicatedfromsortedlist

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	template := "expect %+v, got: %+v"
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	target := []int{1, 2, 5}

	result := deleteDuplicates(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
	}

	input = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	target = []int{2, 3}

	result = deleteDuplicates(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
	}

	input = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	target = []int{1, 2}

	result = deleteDuplicates(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
	}

	input = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	target = []int{}

	result = deleteDuplicates(input)
	if !isSame(target, result) {
		t.Errorf(template, target, result)
	}

}

func isSame(target []int, result *ListNode) bool {
	ind := 0
	node := result

	if len(target) != 0 && result == nil {
		return false
	}

	if len(target) == 0 && result == nil {
		return true
	}

	for node.Next != nil {
		if ind >= len(target) {
			return false
		}

		if target[ind] != node.Val {
			return false
		}

		node = node.Next
		ind++
	}
	return true
}
