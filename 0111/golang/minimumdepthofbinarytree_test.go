package minimumdepthofbinarytree

import (
	"testing"
)

func TestMinDepth(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	target := 2
	result := minDepth(tree)
	if target != result {
		t.Errorf("expect height: %d, got: %d", target, result)
	}
}

func TestMinDepthII(t *testing.T) {
	target := 0
	result := minDepth(nil)
	if target != result {
		t.Errorf("expect height: %d, got: %d", target, result)
	}
}

func TestMinDepthIII(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
	}
	target := 1
	result := minDepth(tree)
	if target != result {
		t.Errorf("expect height: %d, got: %d", target, result)
	}
}

func TestMinDepthIV(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
	}
	target := 2
	result := minDepth(tree)
	if target != result {
		t.Errorf("expect height: %d, got: %d", target, result)
	}
}

func TestMinDepthV(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	target := 2
	result := minDepth(tree)
	if target != result {
		t.Errorf("expect height: %d, got: %d", target, result)
	}
}
