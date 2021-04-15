package palindromelinkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	result := isPalindrome(head)

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeII(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}

	result := isPalindrome(head)

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeIII(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	result := isPalindrome(head)

	target := true
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeIV(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}

	result := isPalindrome(head)

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}

func TestIsPalindromeV(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	result := isPalindrome(head)

	target := false
	if result != target {
		t.Errorf("expect %t, got %t", target, result)
	}
}
