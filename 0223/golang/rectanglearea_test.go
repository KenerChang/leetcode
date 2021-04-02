package rectanglearea

import "testing"

func TestComputeArea(t *testing.T) {
	result := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)

	target := 45
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestComputeAreaII(t *testing.T) {
	result := computeArea(-2, -2, 2, 2, -2, -2, 2, 2)

	target := 16
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestComputeAreaIII(t *testing.T) {
	result := computeArea(-2, -2, 2, 2, -4, -4, -3, -3)

	target := 17
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestComputeAreaIV(t *testing.T) {
	result := computeArea(-2, -2, 2, 2, 3, 3, 4, 4)

	target := 17
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestComputeAreaV(t *testing.T) {
	result := computeArea(-2, -2, 2, 2, -1, -1, 0, 0)

	target := 16
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestComputeAreaVI(t *testing.T) {
	result := computeArea(0, 0, 0, 0, -1, -1, 1, 1)

	target := 4
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestComputeAreaVII(t *testing.T) {
	result := computeArea(-2, -2, 2, 2, -3, -3, 3, -1)

	target := 24
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}

func TestComputeAreaVIII(t *testing.T) {
	result := computeArea(-2, -2, 2, 2, -3, -3, 3, 3)

	target := 36
	if result != target {
		t.Errorf("expect %d, got %d", target, result)
	}
}
