package binarytreepostordertraversal

import (
	"testing"
)

func verify(target, result []int) bool {
	if len(target) != len(result) {
		return false
	}

	for i, numt := range target {
		if numt != result[i] {
			return false
		}
	}
	return true
}

func TestPostorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	target := []int{3, 2, 1}

	result := postorderTraversal(root)
	if !verify(target, result) {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestPostorderTraversalII(t *testing.T) {
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
	target := []int{4, 5, 2, 6, 7, 3, 1}

	result := postorderTraversal(root)
	if !verify(target, result) {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestPostorderTraversalIII(t *testing.T) {
	target := []int{}

	result := postorderTraversal(nil)
	if !verify(target, result) {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestPostorderTraversalIV(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	target := []int{1}

	result := postorderTraversal(root)
	if !verify(target, result) {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestPostorderTraversalV(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	target := []int{2, 1}

	result := postorderTraversal(root)
	if !verify(target, result) {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestPostorderTraversalVI(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	target := []int{2, 1}

	result := postorderTraversal(root)
	if !verify(target, result) {
		t.Errorf("expect %d, got %d", target, result)
	}
}
