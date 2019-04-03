package mergeinterval

import (
	"testing"
)

func TestMerge(t *testing.T) {
	input := []Interval{
		Interval{
			Start: 1, End: 3,
		},
		Interval{
			Start: 2, End: 6,
		},
		Interval{
			Start: 8, End: 10,
		},
		Interval{
			Start: 15, End: 18,
		},
	}
	target := []Interval{
		Interval{
			Start: 1, End: 6,
		},
		Interval{
			Start: 8, End: 10,
		},
		Interval{
			Start: 15, End: 18,
		},
	}
	result := merge(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, but got: %+v", target, result)
	}

	// Input: [[1,4],[4,5]]
	// Output: [[1,5]]
	input = []Interval{
		Interval{
			Start: 1, End: 4,
		},
		Interval{
			Start: 4, End: 5,
		},
	}
	target = []Interval{
		Interval{
			Start: 1, End: 5,
		},
	}
	result = merge(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, but got: %+v", target, result)
	}

	// [[2,3],[4,5],[6,7],[8,9],[1,10]]
	input = []Interval{
		Interval{
			Start: 2, End: 3,
		},
		Interval{
			Start: 4, End: 5,
		},
		Interval{
			Start: 6, End: 7,
		},
		Interval{
			Start: 8, End: 9,
		},
		Interval{
			Start: 1, End: 10,
		},
	}
	target = []Interval{
		Interval{
			Start: 1, End: 10,
		},
	}
	result = merge(input)
	if !isSame(target, result) {
		t.Errorf("expect %+v, but got: %+v", target, result)
	}
}

func isSame(is1, is2 []Interval) bool {
	if len(is1) != len(is2) {
		return false
	}

	same := true
	for idx, interval1 := range is1 {
		interval2 := is2[idx]
		if interval1.Start != interval2.Start || interval1.End != interval2.End {
			same = false
			break
		}
	}
	return same
}
