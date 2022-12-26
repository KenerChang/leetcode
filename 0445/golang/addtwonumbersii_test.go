package addtwonumbersii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbersI(t *testing.T) {
	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
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

	result := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{7, 8, 0, 7}, result.ToInts())
}

func TestAddTwoNumbersII(t *testing.T) {
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

	result := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{8, 0, 7}, result.ToInts())
}

func TestAddTwoNumbersIII(t *testing.T) {
	l1 := &ListNode{
		Val: 0,
	}

	l2 := &ListNode{
		Val: 0,
	}

	result := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{0}, result.ToInts())
}

func TestAddTwoNumbersIV(t *testing.T) {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
		},
	}

	l2 := &ListNode{
		Val: 1,
	}

	result := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{1, 0, 0}, result.ToInts())
}

func TestAddTwoNumbersV(t *testing.T) {
	l1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
		},
	}

	result := addTwoNumbers(l1, l2)

	assert.Equal(t, []int{1, 0, 0}, result.ToInts())
}
