package binarytreelongestconsecutivesequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		root   *TreeNode
		wanted int
	}{
		{
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			3,
		},
		{
			&TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			2,
		},
		{
			&TreeNode{
				Val: 3,
			},
			1,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("root: %+v", c.root), func(t *testing.T) {
			assert.Equal(t, c.wanted, longestConsecutive(c.root))
		})
	}
}
