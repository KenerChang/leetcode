package convertbinarysearchtreetosorteddoublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeToDoublyListI(t *testing.T) {
	root := &Node{Val: 4}
	root.Left = &Node{Val: 2}
	root.Left.Left = &Node{Val: 1}
	root.Left.Right = &Node{Val: 3}
	root.Right = &Node{Val: 5}

	head := treeToDoublyList(root)

	target := []int{1, 2, 3, 4, 5}
	assert.Equal(t, target, head.ValueList())
}

func TestTreeToDoublyListII(t *testing.T) {
	root := &Node{Val: 2}
	root.Left = &Node{Val: 1}
	root.Right = &Node{Val: 3}

	head := treeToDoublyList(root)

	target := []int{1, 2, 3}
	assert.Equal(t, target, head.ValueList())
}

func TestTreeToDoublyListIII(t *testing.T) {
	root := &Node{Val: 2}

	head := treeToDoublyList(root)

	target := []int{2}
	assert.Equal(t, target, head.ValueList())
}

func TestTreeToDoublyListIV(t *testing.T) {
	head := treeToDoublyList(nil)

	assert.Nil(t, head)
}
