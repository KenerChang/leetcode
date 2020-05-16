package pathsum

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}

	target := true
	result := hasPathSum(tree, 22)
	if target != result {
		t.Errorf("expect: %t, got: %t", target, result)
	}
}

func TestPathSumII(t *testing.T) {
	target := false
	result := hasPathSum(nil, 0)
	if target != result {
		t.Errorf("expect: %t, got: %t", target, result)
	}
}

func TestPathSumIII(t *testing.T) {
	target := false
	result := hasPathSum(nil, 22)
	if target != result {
		t.Errorf("expect: %t, got: %t", target, result)
	}
}

func TestPathSumIV(t *testing.T) {
	tree := &TreeNode{
		Val: 0,
	}
	target := true
	result := hasPathSum(tree, 0)
	if target != result {
		t.Errorf("expect: %t, got: %t", target, result)
	}
}

func TestPathSumV(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	target := true
	result := hasPathSum(tree, 15)
	if target != result {
		t.Errorf("expect: %t, got: %t", target, result)
	}
}
