package verticalordertraversalofabinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerticalTraversalI(t *testing.T) {
	root := &TreeNode{
		Val: 3,
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

	assert.Equal(t, [][]int{{9}, {3, 15}, {20}, {7}}, verticalTraversal(root))
}

func TestVerticalTraversalII(t *testing.T) {
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

	assert.Equal(t, [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}, verticalTraversal(root))
}

func TestVerticalTraversalIII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	assert.Equal(t, [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}, verticalTraversal(root))
}
