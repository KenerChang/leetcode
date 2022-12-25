package plusonelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOneI(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	head = plusOne(head)

	assert.Equal(t, []int{1, 2, 4}, head.ToInts())
}

func TestPlusOneII(t *testing.T) {
	head := &ListNode{
		Val: 9,
	}

	head = plusOne(head)

	assert.Equal(t, []int{1, 0}, head.ToInts())
}

func TestPlusOneIII(t *testing.T) {
	head := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	}

	head = plusOne(head)

	assert.Equal(t, []int{1, 0, 0, 0}, head.ToInts())
}

func TestPlusOneIV(t *testing.T) {
	head := &ListNode{
		Val: 0,
	}

	head = plusOne(head)

	assert.Equal(t, []int{1}, head.ToInts())
}

func TestPlusOneV(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
			},
		},
	}

	head = plusOne(head)

	assert.Equal(t, []int{1, 0, 1}, head.ToInts())
}

func TestPlusOneVI(t *testing.T) {
	plusOne(nil)
}
