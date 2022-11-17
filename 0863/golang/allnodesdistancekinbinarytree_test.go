package allnodesdistancekinbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceKI(t *testing.T) {
	target := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	root := &TreeNode{
		Val:  3,
		Left: target,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	assert.ElementsMatch(t, []int{7, 4, 1}, distanceK(root, target, 2))
}

func TestDistanceKII(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}

	assert.ElementsMatch(t, []int{}, distanceK(root, root, 3))
}

func TestDistanceKIII(t *testing.T) {
	target := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	root := &TreeNode{
		Val:   0,
		Right: target,
	}

	assert.ElementsMatch(t, []int{3}, distanceK(root, target, 2))
}

func TestDistanceKIV(t *testing.T) {
	target := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	root := &TreeNode{
		Val:  0,
		Left: target,
	}

	assert.ElementsMatch(t, []int{3}, distanceK(root, target, 2))
}

func TestDistanceKV(t *testing.T) {
	target := &TreeNode{
		Val: 3,
	}
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val:  1,
			Left: target,
		},
	}

	assert.ElementsMatch(t, []int{2}, distanceK(root, target, 3))
}

func TestDistanceKVI(t *testing.T) {
	target := &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 4,
		},
	}
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:   2,
				Right: target,
			},
		},
	}

	assert.ElementsMatch(t, []int{3}, distanceK(root, target, 0))
}
