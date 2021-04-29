package searcha2dmatrixii

import "testing"

func TestSearchMatrix(t *testing.T) {
	martix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	result := searchMatrix(martix, 20)

	target := false
	if result != target {
		t.Errorf("expext %t, got %t", target, result)
	}
}

func TestSearchMatrixII(t *testing.T) {
	martix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	result := searchMatrix(martix, 5)

	target := true
	if result != target {
		t.Errorf("expext %t, got %t", target, result)
	}
}

func TestSearchMatrixIII(t *testing.T) {
	martix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	result := searchMatrix(martix, 1)

	target := true
	if result != target {
		t.Errorf("expext %t, got %t", target, result)
	}
}

func TestSearchMatrixIV(t *testing.T) {
	martix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	result := searchMatrix(martix, 15)

	target := true
	if result != target {
		t.Errorf("expext %t, got %t", target, result)
	}
}

func TestSearchMatrixV(t *testing.T) {
	martix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	result := searchMatrix(martix, 9)

	target := true
	if result != target {
		t.Errorf("expext %t, got %t", target, result)
	}
}

func TestSearchMatrixVI(t *testing.T) {
	martix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	result := searchMatrix(martix, 18)

	target := true
	if result != target {
		t.Errorf("expext %t, got %t", target, result)
	}
}

func TestSearchMatrixVII(t *testing.T) {
	martix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	result := searchMatrix(martix, 30)

	target := true
	if result != target {
		t.Errorf("expext %t, got %t", target, result)
	}
}

func TestSearchMatrixVIII(t *testing.T) {
	martix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	result := searchMatrix(martix, 35)

	target := false
	if result != target {
		t.Errorf("expext %t, got %t", target, result)
	}
}
