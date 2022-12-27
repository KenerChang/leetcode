package swappingnodesinalinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapNodesI(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	head = swapNodes(head, 2)

	assert.Equal(t, []int{1, 4, 3, 2, 5}, head.ToInts())
}

func TestSwapNodesII(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}

	head = swapNodes(head, 1)

	assert.Equal(t, []int{1}, head.ToInts())
}

func TestSwapNodesIII(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	head = swapNodes(head, 5)

	assert.Equal(t, []int{5, 2, 3, 4, 1}, head.ToInts())
}

func TestSwapNodesIV(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	head = swapNodes(head, 3)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, head.ToInts())
}

func TestSwapNodesV(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	head = swapNodes(head, 1)

	assert.Equal(t, []int{5, 2, 3, 4, 1}, head.ToInts())
}

func TestSwapNodesVI(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	head = swapNodes(head, 3)

	assert.Equal(t, []int{1, 2, 4, 3, 5, 6}, head.ToInts())
}
