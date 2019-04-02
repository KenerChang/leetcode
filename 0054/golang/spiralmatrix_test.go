package spiralmatrix

import (
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	input := [][]int{
		[]int{
			1, 2, 3,
		},
		[]int{
			4, 5, 6,
		},
		[]int{
			7, 8, 9,
		},
	}
	target := []int{
		1, 2, 3, 6, 9, 8, 7, 4, 5,
	}

	output := spiralOrder(input)
	if !isSame(target, output) {
		t.Errorf("expect %+v, but got: %+v", target, output)
	}

	input = [][]int{
		[]int{
			1, 2, 3, 4,
		},
		[]int{
			5, 6, 7, 8,
		},
		[]int{
			9, 10, 11, 12,
		},
	}
	target = []int{
		1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7,
	}

	output = spiralOrder(input)
	if !isSame(target, output) {
		t.Errorf("expect %+v, but got: %+v", target, output)
	}

	input = [][]int{
		[]int{
			1, 2,
		},
		[]int{
			3, 4,
		},
	}
	target = []int{1, 2, 4, 3}
	output = spiralOrder(input)
	if !isSame(target, output) {
		t.Errorf("expect %+v, but got: %+v", target, output)
	}

	input = [][]int{
		[]int{
			1, 2,
		},
	}
	target = []int{1, 2}
	output = spiralOrder(input)
	if !isSame(target, output) {
		t.Errorf("expect %+v, but got: %+v", target, output)
	}

	input = [][]int{
		[]int{
			1,
		},
	}
	target = []int{1}
	output = spiralOrder(input)
	if !isSame(target, output) {
		t.Errorf("expect %+v, but got: %+v", target, output)
	}
}

func isSame(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	same := true
	for idx, num1 := range s1 {
		num2 := s2[idx]
		if num1 != num2 {
			same = false
			break
		}
	}
	return same
}
