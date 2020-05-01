package validatebinarysearchtree

import (
	"testing"
)

func TestIsValidBSTI(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	target := true
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTII(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	target := false
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTIII(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
	}

	target := true
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTIV(t *testing.T) {
	target := true
	result := isValidBST(nil)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTV(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}
	target := false
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTVI(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
		},
	}
	target := false
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTVII(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	target := false
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTVIII(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 30,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	target := false
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTVIIII(t *testing.T) {
	tree := &TreeNode{
		Val: -12,
		Right: &TreeNode{
			Val: 85,
			Left: &TreeNode{
				Val: 57,
			},
		},
	}
	target := true
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTX(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 20,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	target := false
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsValidBSTXI(t *testing.T) {
	tree := &TreeNode{
		Val: 2147483647,
	}
	target := true
	result := isValidBST(tree)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}
