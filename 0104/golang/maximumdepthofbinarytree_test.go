package maximumdepthofbinarytree

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
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

	target := 3
	result := maxDepth(tree)

	if target != result {
		t.Errorf("expect level depth: %d, got: %d", target, result)
	}
}

func TestMaxDepthII(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
	}

	target := 1
	result := maxDepth(tree)

	if target != result {
		t.Errorf("expect level depth: %d, got: %d", target, result)
	}
}

func TestMaxDepthIII(t *testing.T) {
	target := 0
	result := maxDepth(nil)

	if target != result {
		t.Errorf("expect level depth: %d, got: %d", target, result)
	}
}

func TestMaxDepthIV(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
	}

	target := 4
	result := maxDepth(tree)

	if target != result {
		t.Errorf("expect level depth: %d, got: %d", target, result)
	}
}
