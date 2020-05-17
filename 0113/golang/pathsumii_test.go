package pathsumii

import (
	"testing"
)

func Verify(s1, s2 [][]int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx, nums1 := range s1 {
		nums2 := s2[idx]
		if len(nums1) != len(nums2) {
			return false
		}

		for j, num1 := range nums1 {
			num2 := nums2[j]
			if num1 != num2 {
				return false
			}
		}
	}
	return true
}

func TestPathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	target := [][]int{
		[]int{5, 4, 11, 2},
		[]int{5, 8, 4, 5},
	}

	result := pathSum(tree, 22)
	if !Verify(target, result) {
		t.Errorf("expect: %+v, got %+v", target, result)
	}
}

func TestPathSumII(t *testing.T) {
	target := [][]int{}

	result := pathSum(nil, 0)
	if !Verify(target, result) {
		t.Errorf("expect: %+v, got %+v", target, result)
	}
}

func TestPathSumIII(t *testing.T) {
	tree := &TreeNode{
		Val: 22,
	}
	target := [][]int{
		[]int{22},
	}

	result := pathSum(tree, 22)
	if !Verify(target, result) {
		t.Errorf("expect: %+v, got %+v", target, result)
	}
}
