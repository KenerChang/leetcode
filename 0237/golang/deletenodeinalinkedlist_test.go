package deletenodeinalinkedlist

import "testing"

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	for i := range target {
		if target[i] != result[i] {
			return false
		}
	}

	return true
}

func TestDeleteNode(t *testing.T) {
	node := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 9,
			},
		},
	}

	deleteNode(node)

	target := []int{1, 9}
	result := node.ToSlice()
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestDeleteNodeII(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
		},
	}

	deleteNode(node)

	target := []int{9}
	result := node.ToSlice()
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestDeleteNodeIII(t *testing.T) {
	node := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
		},
	}

	deleteNode(node)

	target := []int{4}
	result := node.ToSlice()
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestDeleteNodeIV(t *testing.T) {
	node := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},
	}

	deleteNode(node)

	target := []int{1}
	result := node.ToSlice()
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}
