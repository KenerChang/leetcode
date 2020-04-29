package uniquebinarysearchtreesii

import (
	btit "leetcode/0094/golang/binarytreeinordertraversal"
	"testing"
)

func verify(target, result []*btit.TreeNode) bool {
	if len(target) != len(result) {
		return false
	}

	for idx, nt := range target {
		nr := result[idx]

		s1 := btit.InorderTraversal(nt)
		s2 := btit.InorderTraversal(nr)

		if len(s1) != len(s2) {
			return false
		}

		for j, num1 := range s1 {
			num2 := s2[j]

			if num1 != num2 {
				return false
			}
		}
	}
	return true
}

func TestGenerateTrees(t *testing.T) {
	target := []*btit.TreeNode{
		&btit.TreeNode{
			Val: 1,
			Right: &btit.TreeNode{
				Val: 3,
				Left: &btit.TreeNode{
					Val: 2,
				},
			},
		},
		&btit.TreeNode{
			Val: 3,
			Left: &btit.TreeNode{
				Val: 2,
				Left: &btit.TreeNode{
					Val: 1,
				},
			},
		},
		&btit.TreeNode{
			Val: 3,
			Left: &btit.TreeNode{
				Val: 1,
				Right: &btit.TreeNode{
					Val: 2,
				},
			},
		},
		&btit.TreeNode{
			Val: 2,
			Left: &btit.TreeNode{
				Val: 1,
			},
			Right: &btit.TreeNode{
				Val: 3,
			},
		},
		&btit.TreeNode{
			Val: 1,
			Right: &btit.TreeNode{
				Val: 2,
				Right: &btit.TreeNode{
					Val: 3,
				},
			},
		},
	}

	result := generateTrees(3)
	if !verify(target, result) {
		t.Errorf("expect %+v, got: %+v", target, result)
	}
}
