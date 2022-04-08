package closestbinarysearchtreevalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosestValue(t *testing.T) {
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
			Val: 5,
		},
	}

	assert.Equal(t, 4, closestValue(root, 3.714286))
}

func TestClosestValueII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}

	assert.Equal(t, 1, closestValue(root, 3.714286))
}

func TestClosestValueIII(t *testing.T) {
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
			Val: 5,
		},
	}

	assert.Equal(t, 3, closestValue(root, 3.1))
}

func TestClosestValueIV(t *testing.T) {
	root := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 10,
		},
	}

	assert.Equal(t, 4, closestValue(root, 5))
}
