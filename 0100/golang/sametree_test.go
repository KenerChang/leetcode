package sametree

import (
	"testing"
)

func TestIsSameTreeI(t *testing.T) {
	target := true
	result := isSameTree(nil, nil)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsSameTreeII(t *testing.T) {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 23,
		},
	}

	q := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 23,
		},
	}
	target := true
	result := isSameTree(p, q)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsSameTreeIII(t *testing.T) {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 23,
		},
	}

	target := false
	result := isSameTree(p, nil)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsSameTreeIV(t *testing.T) {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 23,
		},
	}

	target := false
	result := isSameTree(nil, p)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}

func TestIsSameTreeV(t *testing.T) {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 23,
		},
	}

	q := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	target := false
	result := isSameTree(p, q)
	if target != result {
		t.Errorf("expect %t, got: %t", target, result)
	}
}
