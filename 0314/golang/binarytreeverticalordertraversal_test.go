package binarytreeverticalordertraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerticalOrderI(t *testing.T) {
	//[3,9,20,null,null,15,7]
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

	assert.Equal(t, [][]int{{9}, {3, 15}, {20}, {7}}, verticalOrder(root))
}

func TestVerticalOrderII(t *testing.T) {
	//[3,9,8,4,0,1,7]
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	assert.Equal(t, [][]int{{4}, {9}, {3, 0, 1}, {8}, {7}}, verticalOrder(root))
}

func TestVerticalOrderIII(t *testing.T) {
	//[3,9,8,4,0,1,7,null,null,null,2,5]
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	assert.Equal(t, [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}}, verticalOrder(root))
}

func TestVerticalOrderIV(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
			},
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

	assert.Equal(t, [][]int{{9, 1}, {3, 5, 15}, {20}, {7}}, verticalOrder(root))
}
