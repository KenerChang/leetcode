package binarytreemaximumpathsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	assert.Equal(t, 6, maxPathSum(root))
}

func TestMaxPathSumII(t *testing.T) {
	root := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val: -2,
		},
		Right: &TreeNode{
			Val: -3,
		},
	}

	assert.Equal(t, -1, maxPathSum(root))
}

func TestMaxPathSumIII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: -3,
		},
	}

	assert.Equal(t, 3, maxPathSum(root))
}

func TestMaxPathSumIV(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: -20,
			Left: &TreeNode{
				Val: -15,
			},
			Right: &TreeNode{
				Val: -7,
			},
		},
	}

	assert.Equal(t, 19, maxPathSum(root))
}

func TestMaxPathSumV(t *testing.T) {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: -5,
			},
			Right: &TreeNode{
				Val: -4,
			},
		},
		Right: &TreeNode{
			Val: -20,
			Left: &TreeNode{
				Val: -15,
			},
			Right: &TreeNode{
				Val: -7,
			},
		},
	}

	assert.Equal(t, 19, maxPathSum(root))
}

func TestMaxPathSumVI(t *testing.T) {
	root := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: -15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	assert.Equal(t, 27, maxPathSum(root))
}

func TestMaxPathSumVII(t *testing.T) {
	root := &TreeNode{
		Val: -10,
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

	assert.Equal(t, 42, maxPathSum(root))
}

func TestMaxPathSumVIII(t *testing.T) {
	root := &TreeNode{
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
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	assert.Equal(t, 48, maxPathSum(root))
}

func TestMaxPathSumIX(t *testing.T) {
	root := &TreeNode{
		Val: -9,
	}

	assert.Equal(t, -9, maxPathSum(root))
}
