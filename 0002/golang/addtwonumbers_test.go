package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbersI(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{7, 0, 8}, l.Ints())
}

func TestAddTwoNumbersII(t *testing.T) {
	l1 := &ListNode{
		Val: 0,
	}

	l2 := &ListNode{
		Val: 0,
	}

	l := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{0}, l.Ints())
}

func TestAddTwoNumbersIII(t *testing.T) {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}

	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}

	l := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{8, 9, 9, 9, 0, 0, 0, 1}, l.Ints())
}

func TestAddTwoNumbersIVs(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
	}

	l2 := &ListNode{
		Val: 0,
	}

	l := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{1}, l.Ints())
}

func TestAddTwoNumbersV(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
	}

	l2 := &ListNode{
		Val: 9,
	}

	l := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{0, 1}, l.Ints())
}
