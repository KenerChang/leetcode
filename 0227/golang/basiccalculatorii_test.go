package basiccalculatorii

import "testing"

func TestCalculate(t *testing.T) {
	result := calculate("3+2*2")

	target := 7
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalculateII(t *testing.T) {
	result := calculate("3+2")

	target := 5
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalculateIII(t *testing.T) {
	result := calculate("1-2-3-4")

	target := -8
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalculateIV(t *testing.T) {
	result := calculate("1-2")

	target := -1
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalculateV(t *testing.T) {
	result := calculate("2/1")

	target := 2
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalculateVI(t *testing.T) {
	result := calculate("1/2")

	target := 0
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalculateVII(t *testing.T) {
	result := calculate("0")

	target := 0
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalculateVIII(t *testing.T) {
	result := calculate("5")

	target := 5
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestCalculateIX(t *testing.T) {
	result := calculate("1+2*5/3+6/4*2")

	target := 6
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
