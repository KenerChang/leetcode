package countunivaluesubtrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountUnivalSubtreesI(t *testing.T) {
	root := TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 5,
			},
			Val: 1,
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Val: 5,
	}

	assert.Equal(t, 4, countUnivalSubtrees(&root))
}

func TestCountUnivalSubtreesII(t *testing.T) {
	assert.Equal(t, 0, countUnivalSubtrees(nil))
}

func TestCountUnivalSubtreesIII(t *testing.T) {
	root := TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 5,
			},
			Val: 5,
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Val: 5,
	}

	assert.Equal(t, 6, countUnivalSubtrees(&root))
}

func TestCountUnivalSubtreesIV(t *testing.T) {
	root := TreeNode{
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
		Val: 1,
	}

	assert.Equal(t, 2, countUnivalSubtrees(&root))
}
