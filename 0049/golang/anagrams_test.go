package anagrams

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	target := [][]string{
		[]string{"ate","eat","tea"},
		[]string{"nat","tan"},
		[]string{"bat"},
	}
	results := groupAnagrams(input)

	if len(target) != len(results) {
		t.Errorf("expect: %+v, got: %+v", target, results)
	}

	input = []string{"eat", "tea", "tan", "ate", "nat", "bat", "batt", "abtt", "tbta"}
	target = [][]string{
		[]string{"ate","eat","tea"},
		[]string{"nat","tan"},
		[]string{"bat"},
		[]string{"batt", "abtt", "tbba"},
	}
	results = groupAnagrams(input)

	if len(target) != len(results) {
		t.Errorf("expect: %+v, got: %+v", target, results)
	}
}