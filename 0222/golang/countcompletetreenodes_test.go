package countcompletetreenodes

import "testing"

func TestCountNodes(t *testing.T) {
	root := &TreeNode{
		Val: 0,
	}

	result := countNodes(root)

	target := 1
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCountNodesII(t *testing.T) {
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

	result := countNodes(root)

	target := 7
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCountNodesIII(t *testing.T) {
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
		},
	}

	result := countNodes(root)

	target := 6
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCountNodesIV(t *testing.T) {
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
		},
	}

	result := countNodes(root)

	target := 5
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCountNodesV(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := countNodes(root)

	target := 4
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
