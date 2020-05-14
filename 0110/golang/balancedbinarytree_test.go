package balancedbinarytree

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
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

	target := true
	result := isBalanced(tree)

	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsBalancedII(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	target := false
	result := isBalanced(tree)

	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}
