package removelinkedlistelements

import (
	"testing"
)

func verify(target, result *ListNode) bool {
	for target != nil && result != nil {
		if target.Val != result.Val {
			return false
		}

		target = target.Next
		result = result.Next
	}

	if target != nil || result != nil {
		return false
	}
	return true
}

func TestRemoveElements(t *testing.T) {
	// 1->2->6->3->4->5->6
	first := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	}
	// 1->2->3->4->5
	target := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	result := removeElements(first, 6)
	if !verify(target, result) {
		t.Errorf("expect %s, got %s", target.String(), result.String())
	}
}
