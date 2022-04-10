package closestbinarysearchtreevalueii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosestKValues(t *testing.T) {
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

	assert.ElementsMatch(t, []int{4, 3}, closestKValues(root, 3.714286, 2))
}

func TestClosestKValuesII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}

	assert.ElementsMatch(t, []int{1}, closestKValues(root, 0, 1))
}

func TestClosestKValuesIII(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}

	assert.ElementsMatch(t, []int{2, 1}, closestKValues(root, 4.142857, 2))
}
