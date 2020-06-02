package palindromepartitioning

import (
	"testing"
)

func verify(target, result [][]string) bool {
	if len(target) != len(result) {
		return false
	}

	for i, row1 := range target {
		if len(row1) != len(result[i]) {
			return false
		}

		for j, str1 := range row1 {
			if str1 != result[i][j] {
				return false
			}
		}
	}
	return true
}

func TestPartition(t *testing.T) {
	target := [][]string{
		[]string{"a", "a", "b"},
		[]string{"aa", "b"},
	}

	result := partition("aab")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionII(t *testing.T) {
	target := [][]string{
		[]string{"a", "a", "b", "a"},
		[]string{"a", "aba"},
		[]string{"aa", "b", "a"},
	}

	result := partition("aaba")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionIII(t *testing.T) {
	target := [][]string{
		[]string{"a", "a", "b", "a", "c"},
		[]string{"a", "aba", "c"},
		[]string{"aa", "b", "a", "c"},
	}

	result := partition("aabac")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionIV(t *testing.T) {
	target := [][]string{
		[]string{"a", "a", "a", "b", "a", "c"},
		[]string{"a", "a", "aba", "c"},
		[]string{"a", "aa", "b", "a", "c"},
		[]string{"aa", "a", "b", "a", "c"},
		[]string{"aa", "aba", "c"},
		[]string{"aaa", "b", "a", "c"},
	}

	result := partition("aaabac")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionV(t *testing.T) {
	target := [][]string{
		[]string{"a", "b", "b", "a", "c"},
		[]string{"a", "bb", "a", "c"},
		[]string{"abba", "c"},
	}

	result := partition("abbac")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionVI(t *testing.T) {
	target := [][]string{
		[]string{"a", "b", "b", "a"},
		[]string{"a", "bb", "a"},
		[]string{"abba"},
	}

	result := partition("abba")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionVII(t *testing.T) {
	target := [][]string{
		[]string{},
	}

	result := partition("")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionVIII(t *testing.T) {
	target := [][]string{
		[]string{"a"},
	}

	result := partition("a")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionIX(t *testing.T) {
	target := [][]string{
		[]string{"a", "b"},
	}

	result := partition("ab")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}

func TestPartitionX(t *testing.T) {
	target := [][]string{
		[]string{"a", "a"},
		[]string{"aa"},
	}

	result := partition("aa")
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}
