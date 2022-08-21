package mergeksortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsMatch(list1, list2 *ListNode) bool {
	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	return list1 == nil && list2 == nil
}

func TestMergeKListsI(t *testing.T) {
	input := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	}

	target := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}}

	assert.True(t, IsMatch(mergeKLists(input), target))
}

func TestMergeKListsII(t *testing.T) {
	input := []*ListNode{}

	assert.True(t, IsMatch(mergeKLists(input), nil))
}

func TestMergeKListsIII(t *testing.T) {
	input := []*ListNode{
		nil,
	}

	assert.True(t, IsMatch(mergeKLists(input), nil))
}
