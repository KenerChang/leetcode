package binarytreerightsideview

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

func TestRightSideView(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	target := []int{1, 3, 4}

	result := rightSideView(root)
	if !verify(target, result) {
		t.Errorf("exepct %v, got %v", target, result)
	}
}

func TestRightSideViewII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	target := []int{1, 3, 4}

	result := rightSideView(root)
	if !verify(target, result) {
		t.Errorf("exepct %v, got %v", target, result)
	}
}

func TestRightSideViewIII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	target := []int{1, 3, 4}

	result := rightSideView(root)
	if !verify(target, result) {
		t.Errorf("exepct %v, got %v", target, result)
	}
}

func TestRightSideViewIV(t *testing.T) {
	target := []int{}

	result := rightSideView(nil)
	if !verify(target, result) {
		t.Errorf("exepct %v, got %v", target, result)
	}
}

func TestRightSideViewV(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	target := []int{1, 2}

	result := rightSideView(root)
	if !verify(target, result) {
		t.Errorf("exepct %v, got %v", target, result)
	}
}

func TestRightSideViewVI(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	target := []int{1, 3, 5}

	result := rightSideView(root)
	if !verify(target, result) {
		t.Errorf("exepct %v, got %v", target, result)
	}
}
