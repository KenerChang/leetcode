package sumroottoleafnumbers

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	target := 25
	result := sumNumbers(tree)
	if target != result {
		t.Errorf("expect: %d, got: %d", target, result)
	}
}
func TestSumNumbersII(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	target := 1026
	result := sumNumbers(tree)
	if target != result {
		t.Errorf("expect: %d, got: %d", target, result)
	}
}

func TestSumNumbersIII(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
	}

	target := 4
	result := sumNumbers(tree)
	if target != result {
		t.Errorf("expect: %d, got: %d", target, result)
	}
}
