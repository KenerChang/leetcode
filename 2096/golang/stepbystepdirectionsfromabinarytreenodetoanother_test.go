package stepbystepdirectionsfromabinarytreenodetoanother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirectionsI(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	assert.Equal(t, "UURL", getDirections(root, 3, 6))
}

func TestGetDirectionsII(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}

	assert.Equal(t, "L", getDirections(root, 2, 1))
}

func TestGetDirectionsIII(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}

	assert.Equal(t, "U", getDirections(root, 1, 2))
}

func TestGetDirectionsIV(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 1,
		},
	}

	assert.Equal(t, "R", getDirections(root, 2, 1))
}

func TestGetDirectionsV(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 12,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 7,
						},
					},
					Right: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			Right: &TreeNode{
				Val: 13,
				Right: &TreeNode{
					Val: 15,
					Right: &TreeNode{
						Val: 14,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, "UURR", getDirections(root, 6, 15))
}
