package evaluatereversepolishnotation

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	input := []string{"2", "1", "+", "3", "*"}
	target := 9

	result := evalRPN(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestEvalRPNII(t *testing.T) {
	input := []string{"4", "13", "5", "/", "+"}
	target := 6

	result := evalRPN(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestEvalRPNIII(t *testing.T) {
	input := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	target := 22

	result := evalRPN(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestEvalRPNIV(t *testing.T) {
	input := []string{"1", "2", "3", "*", "5", "+", "-"}
	target := -10

	result := evalRPN(input)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
