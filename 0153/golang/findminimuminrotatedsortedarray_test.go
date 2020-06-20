package findminimuminrotatedsortedarray

import (
	"testing"
)

// func TestFindMin(t *testing.T) {
// 	nums := []int{3, 4, 5, 1, 2}
// 	target := 1

// 	result := findMin(nums)
// 	if target != result {
// 		t.Errorf("expect %d, got %d", target, result)
// 	}
// }

// func TestFindMinII(t *testing.T) {
// 	nums := []int{4, 5, 6, 7, 0, 1, 2}
// 	target := 0

// 	result := findMin(nums)
// 	if target != result {
// 		t.Errorf("expect %d, got %d", target, result)
// 	}
// }

// func TestFindMinIII(t *testing.T) {
// 	nums := []int{4}
// 	target := 4

// 	result := findMin(nums)
// 	if target != result {
// 		t.Errorf("expect %d, got %d", target, result)
// 	}
// }

// func TestFindMinIV(t *testing.T) {
// 	nums := []int{5, 4}
// 	target := 4

// 	result := findMin(nums)
// 	if target != result {
// 		t.Errorf("expect %d, got %d", target, result)
// 	}
// }

// func TestFindMinV(t *testing.T) {
// 	nums := []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
// 	target := -5

// 	result := findMin(nums)
// 	if target != result {
// 		t.Errorf("expect %d, got %d", target, result)
// 	}
// }

func TestFindMinVI(t *testing.T) {
	nums := []int{57, 58, 59, 62, 63, 66, 68, 72, 73, 74, 75, 76, 77, 78, 80, 81, 86, 95, 96, 97, 98, 100, 101, 102, 103, 110, 119, 120, 121, 123, 125, 126, 127, 132, 136, 144, 145, 148, 149, 151, 152, 160, 161, 163, 166, 168, 169, 170, 173, 174, 175, 178, 182, 188, 189, 192, 193, 196, 198, 199, 200, 201, 202, 212, 218, 219, 220, 224, 225, 229, 231, 232, 234, 237, 238, 242, 248, 249, 250, 252, 253, 254, 255, 257, 260, 266, 268, 270, 273, 276, 280, 281, 283, 288, 290, 291, 292, 294, 295, 298, 299, 4, 10, 13, 15, 16, 17, 18, 20, 22, 25, 26, 27, 30, 31, 34, 38, 39, 40, 47, 53, 54}
	target := 4

	result := findMin(nums)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
