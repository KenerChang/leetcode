package findleavesofbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLeaves(t *testing.T) {
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

	assert.Equal(t, [][]int{{4, 5, 3}, {2}, {1}}, findLeaves(root))
}

func TestFindLeavesII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}

	assert.Equal(t, [][]int{{1}}, findLeaves(root))
}
