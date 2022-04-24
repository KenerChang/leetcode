package recoverbinarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func toArray(root *TreeNode, results []int) []int {
	if root == nil {
		return results
	}

	if root.Left == nil && root.Right == nil {
		results = append(results, root.Val)
		return results
	}

	results = toArray(root.Left, results)
	results = append(results, root.Val)
	results = toArray(root.Right, results)

	// fmt.Printf("results: %+v\n", results)
	return results
}

func TestRecoverTreeI(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	recoverTree(root)

	assert.Equal(t, []int{1, 2, 3}, toArray(root, []int{}))
}

func TestRecoverTreeII(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	recoverTree(root)

	assert.Equal(t, []int{1, 2, 3, 4}, toArray(root, []int{}))
}
