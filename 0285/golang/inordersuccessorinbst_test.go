package inordersuccessorinbst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	cases := []struct {
		root   *TreeNode
		p      *TreeNode
		wanted *TreeNode
	}{
		{&TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		}, &TreeNode{Val: 2}, &TreeNode{Val: 3}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("root: %d", c.root.Val), func(t *testing.T) {
			result := inorderSuccessor(c.root, c.p)

			if c.wanted == nil || result == nil {
				assert.Equal(t, c.wanted, result)
			} else {
				assert.Equal(t, c.wanted.Val, result.Val)
			}
		})
	}
}
