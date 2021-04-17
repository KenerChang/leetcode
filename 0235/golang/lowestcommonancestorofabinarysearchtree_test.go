package lowestcommonancestorofabinarysearchtree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 4})

	target := &TreeNode{
		Val: 2,
	}

	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorII(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 8})

	target := &TreeNode{
		Val: 6,
	}

	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorIII(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 8}, &TreeNode{Val: 9})

	target := &TreeNode{
		Val: 8,
	}

	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorIV(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 8}, &TreeNode{Val: 2})

	target := &TreeNode{
		Val: 6,
	}

	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorV(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 9}, &TreeNode{Val: 8})

	target := &TreeNode{
		Val: 8,
	}

	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}

func TestLowestCommonAncestorVI(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}

	result := lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 1})

	target := &TreeNode{
		Val: 2,
	}

	if result.Val != target.Val {
		t.Errorf("expect %d, got %d", target.Val, result.Val)
	}
}
