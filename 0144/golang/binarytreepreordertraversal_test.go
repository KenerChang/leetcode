package binarytreepreordertraversal

import (
	"testing"
)

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	for i, num1 := range target {
		if num1 != result[i] {
			return false
		}
	}
	return true
}

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	target := []int{1, 2, 3}
	result := preorderTraversal(root)

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestPreorderTraversalIII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	target := []int{1, 2, 4, 5, 3, 6, 7}
	result := preorderTraversal(root)

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
