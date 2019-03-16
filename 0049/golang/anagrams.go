package anagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	// sort each str alphabetically so that anagrams have same expression
	angramMap := map[string][]string{}
	for _, str := range strs {
		sortedStr := sortString(str)
		if _, ok := angramMap[sortedStr]; ok {
			angramMap[sortedStr] = append(angramMap[sortedStr], str)
		} else {
			angramMap[sortedStr] = []string{str}
		}
	}

	results := [][]string{}
	for _, strs := range angramMap {
		results = append(results, strs)
	}
	return results
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func sortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

