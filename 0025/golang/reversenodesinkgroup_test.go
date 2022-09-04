package reversenodesinkgroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	expected := []int{2, 1, 4, 3, 5}

	actual := reverseKGroup(input, 2).Vals()
	assert.Equal(t, expected, actual)
}

func TestReverseKGroupII(t *testing.T) {
	input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	expected := []int{1, 2, 3, 4, 5}

	actual := reverseKGroup(input, 1).Vals()
	assert.Equal(t, expected, actual)
}

func TestReverseKGroupIII(t *testing.T) {
	input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	expected := []int{3, 2, 1, 4, 5}

	actual := reverseKGroup(input, 3).Vals()
	assert.Equal(t, expected, actual)
}

func TestReverseKGroupIV(t *testing.T) {
	input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	expected := []int{4, 3, 2, 1, 5}

	actual := reverseKGroup(input, 4).Vals()
	assert.Equal(t, expected, actual)
}

func TestReverseKGroupV(t *testing.T) {
	input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	expected := []int{5, 4, 3, 2, 1}

	actual := reverseKGroup(input, 5).Vals()
	assert.Equal(t, expected, actual)
}

func TestReverseKGroupVI(t *testing.T) {
	input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}

	expected := []int{2, 1, 4, 3, 6, 5}

	actual := reverseKGroup(input, 2).Vals()
	assert.Equal(t, expected, actual)
}
