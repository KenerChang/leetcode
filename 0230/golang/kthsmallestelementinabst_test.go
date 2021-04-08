package kthsmallestelementinabst

import "testing"

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	result := kthSmallest(root, 1)
	target := 1

	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestKthSmallestII(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	result := kthSmallest(root, 4)
	target := 4

	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestKthSmallestIII(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	result := kthSmallest(root, 5)
	target := 5

	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}
