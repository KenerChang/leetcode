package lowestcommonancestorofabinarytree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 1})

	target := &TreeNode{Val: 3}
	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorII(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 4})

	target := &TreeNode{Val: 5}
	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorIII(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 1}, &TreeNode{Val: 2})

	target := &TreeNode{Val: 2}
	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorIV(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 1})

	target := &TreeNode{Val: 2}
	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorV(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 4}, &TreeNode{Val: 5})

	target := &TreeNode{Val: 5}
	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}
