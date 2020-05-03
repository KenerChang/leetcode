package symmetrictree

import (
	"testing"
)

func TestIsSymmetricI(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 10,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	target := false
	result := isSymmetric(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsSymmetricII(t *testing.T) {
	target := true
	result := isSymmetric(nil)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsSymmetricIII(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	target := false
	result := isSymmetric(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsSymmetricIV(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	target := false
	result := isSymmetric(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsSymmetricV(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	target := true
	result := isSymmetric(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}
