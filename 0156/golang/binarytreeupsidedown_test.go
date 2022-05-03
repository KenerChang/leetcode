package binarytreeupsidedown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func preorder(root *TreeNode, results []int) []int {
	if root == nil {
		results = append(results, -1)
		return results
	}

	results = append(results, root.Val)

	results = preorder(root.Left, results)
	results = preorder(root.Right, results)

	return results
}

func TestUpsideDownBinaryTree(t *testing.T) {
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
	result := upsideDownBinaryTree(root)

	target := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 5,
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

	assert.Equal(t, preorder(target, []int{}), preorder(result, []int{}))
}

func TestUpsideDownBinaryTreeII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	result := upsideDownBinaryTree(root)

	target := &TreeNode{
		Val: 1,
	}

	assert.Equal(t, preorder(target, []int{}), preorder(result, []int{}))
}

func TestUpsideDownBinaryTreeIII(t *testing.T) {
	result := upsideDownBinaryTree(nil)

	assert.Equal(t, preorder(nil, []int{}), preorder(result, []int{}))
}

func TestUpsideDownBinaryTreeIV(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	result := upsideDownBinaryTree(root)

	target := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	assert.Equal(t, preorder(target, []int{}), preorder(result, []int{}))
}

func TestUpsideDownBinaryTreeV(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	result := upsideDownBinaryTree(root)

	target := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 1,
		},
	}

	assert.Equal(t, preorder(target, []int{}), preorder(result, []int{}))
}
