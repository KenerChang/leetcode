package rectanglearea

import (
	"math"
)

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	areaA := (C - A) * (D - B)
	areaB := (G - E) * (H - F)

	if areaA == 0 || areaB == 0 {
		return areaA + areaB
	}

	if coverArea, coverd := isCoverd(A, B, C, D, E, F, G, H); coverd {
		return areaA + areaB - coverArea
	}
	return areaA + areaB
}

func isCoverd(A int, B int, C int, D int, E int, F int, G int, H int) (int, bool) {
	// not covered
	if A >= G || C <= E || B >= H || D <= F {
		return 0, false
	}

	width := min(G-E, C-E, G-A, C-A)
	height := min(D-F, H-F, H-B, D-B)
	return width * height, true
}

func min(nums ...int) int {
	result := math.MaxInt32

	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return result
}
