package summaryranges

import "testing"

func verify(target, result []string) bool {
	if len(target) != len(result) {
		return false
	}

	for i, r := range target {
		if result[i] != r {
			return false
		}
	}
	return true
}

func TestSummaryRanges(t *testing.T) {
	result := summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})

	target := []string{"0", "2->4", "6", "8->9"}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestSummaryRangesII(t *testing.T) {
	result := summaryRanges([]int{0, 1, 2, 4, 5, 7})

	target := []string{"0->2", "4->5", "7"}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestSummaryRangesIII(t *testing.T) {
	result := summaryRanges([]int{})

	target := []string{}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestSummaryRangesIV(t *testing.T) {
	result := summaryRanges([]int{-1})

	target := []string{"-1"}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestSummaryRangesV(t *testing.T) {
	result := summaryRanges([]int{0, 1, 2, 3, 4})

	target := []string{"0->4"}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}

func TestSummaryRangesVI(t *testing.T) {
	result := summaryRanges([]int{1, 3, 5, 7, 9})

	target := []string{"1", "3", "5", "7", "9"}

	if !verify(target, result) {
		t.Errorf("expect %v, got %v", target, result)
	}
}
