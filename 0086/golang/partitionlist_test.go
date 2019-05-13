package partitionlist

import (
	"testing"
)

func TestPartition(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	target := []int{1, 2, 2, 4, 3, 5}

	result := partition(input, 3)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 3,
									Next: &ListNode{
										Val:  1,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	target = []int{2, 2, 1, 2, 1, 4, 3, 5, 3}

	result = partition(input, 3)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  2,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	target = []int{1, 2, 2, 4, 3, 6, 5}

	result = partition(input, 3)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = nil
	target = []int{}

	result = partition(input, 0)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	target = []int{1, 2}

	result = partition(input, 2)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	target = []int{1, 1, 2}

	result = partition(input, 2)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	target = []int{1, 2, 3}

	result = partition(input, 2)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	target = []int{1, 2, 3}

	result = partition(input, 2)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val:  1,
		Next: nil,
	}
	target = []int{1}

	result = partition(input, 1)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	target = []int{1, 1}

	result = partition(input, 1)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	target = []int{1, 3}

	result = partition(input, 2)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	target = []int{3, 1}

	result = partition(input, 4)
	if !isSame(target, result) {
		errorf(t, target, result)
	}

	input = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	target = []int{1, 2, 3}

	result = partition(input, 3)
	if !isSame(target, result) {
		errorf(t, target, result)
	}
}

func isSame(target []int, result *ListNode) bool {
	ind := 0
	node := result

	if len(target) != 0 && result == nil {
		return false
	}

	if len(target) == 0 && result == nil {
		return true
	}

	for node != nil {
		if ind >= len(target) {
			return false
		}

		if target[ind] != node.Val {
			return false
		}

		node = node.Next
		ind++
	}
	if ind != len(target) {
		return false
	}
	return true
}

func errorf(t *testing.T, target []int, result *ListNode) {
	template := "expect %+v, got: %+v"

	exist := map[*ListNode]bool{}
	cNode := result
	resultVals := []int{}
	for cNode != nil {
		if _, ok := exist[cNode]; ok {
			break
		}
		resultVals = append(resultVals, cNode.Val)
		exist[cNode] = true
		cNode = cNode.Next
	}
	t.Errorf(template, target, resultVals)
}
