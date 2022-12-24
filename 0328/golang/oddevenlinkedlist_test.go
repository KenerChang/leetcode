package oddevenlinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddEvenListI(t *testing.T) {
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

	oddEvenList(head)

	assert.Equal(t, []int{1, 3, 5, 2, 4}, head.toArray())
}

func TestOddEvenListII(t *testing.T) {
	head := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 7,
							},
						},
					},
				},
			},
		},
	}

	oddEvenList(head)

	assert.Equal(t, []int{2, 3, 6, 7, 1, 5, 4}, head.toArray())
}

func TestOddEvenListIII(t *testing.T) {
	head := &ListNode{
		Val: 2,
	}

	oddEvenList(head)

	assert.Equal(t, []int{2}, head.toArray())
}

func TestOddEvenListIV(t *testing.T) {
	oddEvenList(nil)
}

func TestOddEvenListV(t *testing.T) {
	head := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}

	oddEvenList(head)

	assert.Equal(t, []int{2, 4}, head.toArray())
}
