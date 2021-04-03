package invertbinarytree

import "testing"

func verify(treeA, treeB *TreeNode) bool {
	if treeA == nil && treeB == nil {
		return true
	}

	if treeA == nil || treeB == nil {
		return false
	}

	if treeA.Val != treeB.Val {
		return false
	}

	return verify(treeA.Left, treeB.Left) && verify(treeA.Right, treeB.Right)
}

func TestInvertTree(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	result := invertTree(root)

	target := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestInvertTreeII(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := invertTree(root)

	target := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestInvertTreeIII(t *testing.T) {
	root := &TreeNode{
		Val: 2,
	}

	result := invertTree(root)

	target := &TreeNode{
		Val: 2,
	}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestInvertTreeIV(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}

	result := invertTree(root)

	target := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 1,
		},
	}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestInvertTreeV(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 1,
		},
	}

	result := invertTree(root)

	target := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
